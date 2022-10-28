# STIR/SHAKEN Certificate Repository Compliance

## connexcs.com

Code: w_atis_pem_certificate_chain\
Source: ATIS-1000080\
Description: ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain

### https://app.connexcs.com/api/stir-shaken/cert/41.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| w_atis_pem_certificate_chain | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |
| e_http_status_200 | error | HTTP | HTTP response shall have StatusCode 200, but it is 503 Service Unavailable |


Generated: 28/10/2022 at 19:21:41