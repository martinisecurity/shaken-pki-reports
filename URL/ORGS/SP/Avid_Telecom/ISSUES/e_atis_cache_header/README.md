# STIR/SHAKEN Certificate Repository Compliance

## Avid Telecom

Code: e_atis_cache_header\
Source: ATIS-1000074\
Description: The STI-VS shall implement the cache behavior described in RFC7234. If the HTTP response does not include any recognized caching directives or indicates caching for less than 24 hours, then the STI-VS should cache the HTTP response for 24 hours

### https://ssc.getsipnav.com/certs/d29187037eabb3b60e24baba5c0b3dc424e004b9

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |


Generated: 28/10/2022 at 18:15:07