package auth_client

import (
	"encoding/json"
	"fmt"
)

// LoginOutcome is what the browser could observe after a login attempt.
//
// It carries EVIDENCE, never a verdict. Whether a given outcome means "wrong
// password", "two-factor prompt" or "bot challenge" is a policy question, and
// the vocabulary for it belongs to the caller — this package cannot know which
// classes a caller distinguishes or what it does with each. Deciding here would
// also freeze the mapping behind a module bump, which is the wrong place for a
// rule that has to be corrected from live observation.
//
// The reason any of this exists: a login attempt that produces no session
// cookie has several distinct causes that are indistinguishable from the
// cookie's absence alone. A rejected password, a 2FA prompt and a Cloudflare
// interstitial all end with no FX_RM, so inferring the cause from the missing
// cookie collapses three problems — one of which the user can fix, one of which
// they cannot, and one of which is an infrastructure signal — into a single
// unactionable class.
type LoginOutcome struct {
	// FinalURL is where the browser ended up after submitting the form.
	// A login that Fantrax accepted navigates away from the login page; one it
	// rejected does not.
	FinalURL string `json:"url"`

	// Title is document.title. Cloudflare's interstitial is identifiable by it
	// ("Just a moment...") even when its DOM markers change.
	Title string `json:"title"`

	// Matched names the probes whose selector was present in the DOM, in the
	// order the caller supplied them.
	Matched []string `json:"-"`

	// Texts holds the trimmed visible text of each matched probe that asked for
	// it, keyed by probe name. This is where a form's own error message lands
	// ("Invalid username or password"), which is the difference between telling
	// a user their password is wrong and telling them something unspecified
	// went wrong.
	Texts map[string]string `json:"-"`

	// Raw is the unparsed evidence object, kept so a caller can log the whole
	// thing. A class that was guessed rather than measured is worth exactly as
	// much as the evidence someone can go back and read.
	Raw map[string]any `json:"-"`
}

// Has reports whether the named probe matched.
func (o *LoginOutcome) Has(name string) bool {
	if o == nil {
		return false
	}
	for _, m := range o.Matched {
		if m == name {
			return true
		}
	}
	return false
}

// Text returns the captured text for a probe, or "" if it did not match or did
// not capture.
func (o *LoginOutcome) Text(name string) string {
	if o == nil {
		return ""
	}
	return o.Texts[name]
}

// LoginProbe is a named CSS selector to look for after the login attempt.
//
// Probes are supplied by the caller rather than hardcoded so that the selector
// map can be corrected without a module bump. Fantrax's login page is an Angular
// app whose markup is not a contract, so any selector here is a hypothesis with
// a shelf life — and a hypothesis that can only be revised by publishing a new
// version of a library is one that will not be revised.
type LoginProbe struct {
	// Name is the caller's label, echoed back in LoginOutcome.Matched.
	Name string
	// Selector is a CSS selector passed to document.querySelector.
	Selector string
	// CaptureText records the element's visible text in LoginOutcome.Texts.
	CaptureText bool
}

// DefaultLoginProbes is a starting hypothesis, not a measured selector map.
//
// Every entry here is a guess at markup this package does not control, and the
// point of capturing FinalURL, Title and LoginFormPresent alongside them is that
// those three are informative even when every probe misses. A caller should log
// the whole outcome on failure and tighten this list from what real attempts
// actually produce.
var DefaultLoginProbes = []LoginProbe{
	// The login form itself. Still present => the browser never left the login
	// page, which is the single most discriminating fact available: Fantrax
	// accepting a login navigates away.
	{Name: "login_form", Selector: `input[formcontrolname="email"]`},

	// Angular Material renders field- and form-level errors as <mat-error>.
	{Name: "form_error", Selector: `mat-error`, CaptureText: true},
	{Name: "error_text", Selector: `.error-message, .login-error, [class*="error"]`, CaptureText: true},

	// A one-time-code field means credentials were ACCEPTED and a second factor
	// is being demanded — the user's account is fine and no amount of retrying
	// the password will help.
	{Name: "otp", Selector: `input[autocomplete="one-time-code"], input[formcontrolname="code"], input[name*="otp" i]`},

	// Cloudflare. Distinct from 2FA because it is our problem, not the user's:
	// it means the headless browser is being treated as a bot.
	{Name: "cloudflare", Selector: `#challenge-form, #cf-challenge-running, [id^="cf-chl"], iframe[src*="challenges.cloudflare.com"]`},
}

// loginEvidenceJS builds a single expression that collects every fact in one
// round trip. One evaluation rather than a WaitVisible per probe matters
// because WaitVisible BLOCKS until its timeout when an element is absent, and
// absence is the normal case for most of these — probing serially would spend a
// timeout per miss on exactly the failing path that most needs to answer
// quickly.
func loginEvidenceJS(probes []LoginProbe) (string, error) {
	spec, err := json.Marshal(probes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`(() => {
  const probes = %s;
  const out = { url: location.href, title: document.title, matched: [], texts: {} };
  for (const p of probes) {
    let el = null;
    try { el = document.querySelector(p.Selector); } catch (e) { continue; }
    if (!el) continue;
    out.matched.push(p.Name);
    if (p.CaptureText) {
      const t = (el.innerText || el.textContent || '').trim();
      if (t) out.texts[p.Name] = t.slice(0, 500);
    }
  }
  return out;
})()`, spec), nil
}
