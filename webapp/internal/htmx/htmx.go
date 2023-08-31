package htmx

import "net/http"

const (
	// HXBoosted indicates that the request is via an element using hx-boost.
	// See https://htmx.org/reference/#request_headers
	HXBoosted = "HX-Boosted"

	// HXCurrentURL is the current URL of the browser
	// See https://htmx.org/reference/#request_headers
	HXCurrentURL = "HX-Current-URL"

	// HXHistoryResoreRequest is true if the request is for history
	// restoration after a miss in the local history cache
	// See https://htmx.org/reference/#request_headers
	HXHistoryResoreRequest = "HX-History-Restore-Request"

	// HXPrompt the user response to an hx-prompt
	// See https://htmx.org/reference/#request_headers
	HXPrompt = "HX-Prompt"

	// HXRequest is always true for requests from htmx
	// See https://htmx.org/reference/#request_headers
	HXRequest = "HX-Request"

	// HXTarget is the id of the target element if it exists
	// See https://htmx.org/reference/#request_headers
	HXTarget = "HX-Target"

	// HXTriggerName is the name of the triggered element if it exists
	// See https://htmx.org/reference/#request_headers
	HXTriggerName = "HX-Trigger-Name"

	// HXTriggerID is the id of the triggered element if it exists
	// See https://htmx.org/reference/#request_headers
	HXTriggerID = "HX-Trigger"

	// HXLocation allows you to do a client-side redirect that does not do a full page reload
	// See https://htmx.org/reference/#response_headers
	HXLocation = "HX-Location"

	// HXPushURL pushes a new url into the history stack
	// See https://htmx.org/reference/#response_headers
	HXPushURL = "HX-Push-Url"

	// HXRedirect can be used to do a client-side redirect to a new location
	// See https://htmx.org/reference/#response_headers
	HXRedirect = "HX-Redirect"

	// HXRefresh triggers a full client side refresh if set to “true”
	// See https://htmx.org/reference/#response_headers
	HXRefresh = "HX-Refresh"

	// HXReplaceURL replaces the current URL in the location bar
	// See https://htmx.org/reference/#response_headers
	HXReplaceURL = "HX-Replace-Url"

	// HXReswap allows you to specify how the response will be swapped. See
	// hx-swap (https://htmx.org/attributes/hx-swap/) for possible values
	// See https://htmx.org/reference/#response_headers
	HXReswap = "HX-Reswap"

	// HXRetarget is a CSS selector that updates the target of the content
	// update to a different element on the page
	// See https://htmx.org/reference/#response_headers
	HXRetarget = "HX-Retarget"

	// HXReselect is a CSS selector that allows you to choose which part of
	// the response is used to be swapped in. Overrides an existing
	// hx-select on the triggering element
	// See https://htmx.org/reference/#response_headers
	HXReselect = "HX-Reselect"

	// HXTrigger allows you to trigger client side events, see the
	// documentation for more info
	//
	// Note: See HXTriggerID for the "HX-Trigger" request header
	// See https://htmx.org/reference/#response_headers
	HXTrigger = "HX-Trigger"

	// HXTriggerAfterSettle allows you to trigger client side events, see
	// the documentation for more info
	// See https://htmx.org/reference/#response_headers
	HXTriggerAfterSettle = "HX-Trigger-After-Settle"

	// HXTriggerAfterSwap allows you to trigger client side events, see the
	// documentation for more info
	// See https://htmx.org/reference/#response_headers
	HXTriggerAfterSwap = "HX-Trigger-After-Swap"
)

// Redirect replies to the request with a redirect url. The "HX-Redirect"
// header is used for htmx reuests. Other requests will get a plain http
// redirect. See http.Redirect.
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	if r.Header.Get(HXRequest) != "true" {
		http.Redirect(w, r, url, code)
		return
	}

	w.Header().Set(HXRedirect, url)
}
