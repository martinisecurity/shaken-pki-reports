# STIR/SHAKEN Certificate Repository Compliance

## Drop Inc

Code: w_atis_protocol\
Source: ATIS-1000080\
Description: The verifier should not dereference any protocol other than https or a port other than 443 or 8443

### http://sti.comsapi.com/258k/ca.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_protocol | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |


Generated: 28/10/2022 at 18:54:21