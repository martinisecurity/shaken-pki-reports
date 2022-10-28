# STIR/SHAKEN Certificate Repository Compliance

## Coztel Carrier

Code: e_atis_cache_header\
Source: ATIS-1000074\
Description: The STI-VS shall implement the cache behavior described in RFC7234. If the HTTP response does not include any recognized caching directives or indicates caching for less than 24 hours, then the STI-VS should cache the HTTP response for 24 hours

### http://5.161.95.22/191c4c42dd7fa6115e84100637e42c99.cer

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| w_atis_protocol | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |


Generated: 28/10/2022 at 18:15:07