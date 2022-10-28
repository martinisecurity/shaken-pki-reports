# STIR/SHAKEN Certificate Repository Compliance

## connexcs.com

| Code | Source | Instances |
|------|--------|-----------|
| [e_atis_cache_header](ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 | 1 |
| [w_atis_content_type](ISSUES/w_atis_content_type/README.md) | ATIS-1000080 | 1 |
| [w_aits_pem_certificate_chain](ISSUES/w_aits_pem_certificate_chain/README.md) | ATIS-1000080 | 1 |
| [e_http_status_200](ISSUES/e_http_status_200/README.md) | HTTP | 1 |

### https://app.connexcs.com/api/stir-shaken/cert/41.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_http_status_200 | error | HTTP | HTTP response shall have StatusCode 200, but it is 503 Service Unavailable |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| w_aits_pem_certificate_chain | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |


Generated: 28/10/2022 at 18:15:07