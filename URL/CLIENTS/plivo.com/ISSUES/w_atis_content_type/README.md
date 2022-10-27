# STIR/SHAKEN Certificate Repository Compliance

## plivo.com

Code: w_atis_content_type\
Source: ATIS-1000080\
Description: ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain
### https://certificate.zt.plivo.com/cert09062023.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |


Generated: 27/10/2022 at 21:42:12