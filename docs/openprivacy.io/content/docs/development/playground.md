---
title: "Playground"
description: "OPV Playground"
images: []
menu:
  docs:
    parent: "development"
weight: 1
toc: true
---

{{< alert icon="👉" text="All the testing grant tokens or API tokens used below are for demo purpose only." >}}

## Data Plane Demo

Playground demo of tokenization and detokenization via the Data Plane: `https://playground.openprivacy.io`.

- For tokenization, we are using a public OPV grant token that's only writtable to data plane's `/js/v1/facts` route.
- For detokenization, we are using a private OPV grant token that's only readable to data plane's `/api/v1/facts/*` route. Note that we highly recommend one **don't expose** the token with read permissions to the public JavaScript client. This detokenization is for demo purposes only.

<iframe width="100%" height="300" src="//jsfiddle.net/vb53eakL/embedded/result,js,html,css/dark/" allowfullscreen="allowfullscreen" allowpaymentrequest frameborder="0"></iframe>

----

## Proxy Plane Demo

Playground demo of sending API requests to Checkr via the Proxy Plane: `proxy-playground.openprivacy.io`.

- The definition of the proxy routing can be found at [opv-proxyplane-http.example.json](https://github.com/roney492/opv/blob/53eb70c1ce9aaaa897863982efb468df487ce7c0/cmd/proxyplane/opv-proxyplane-http.example.json#L105).
- We simulated that once the internal system have tokenzied `facts` (i.e. sensitive PIIs), it can talks to a dedicated proxy plane route for automated detokenization when sending the requests to external vendors like [Checkr API](https://api.checkr.com).
- You can also inspect the network requests directly from your browser (tldr - press `F12`) to check the actual payload.

<iframe width="100%" height="600" src="//jsfiddle.net/682vwdtu/2/embedded/result,js,html,css/dark/" allowfullscreen="allowfullscreen" allowpaymentrequest frameborder="0"></iframe>
