# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Code: w_atis_content_type\
Source: ATIS-1000080\
Description: ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain
### https://app.connexcs.com/api/stir-shaken/cert/41.crt

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_aits_pem_certificate_chain | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |
| e_http_status_200 | error | HTTP | HTTP response shall have StatusCode 200, but it is 503 Service Unavailable |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

### https://cr.sansay.com/Mitel_Cloud_Services_Inc._670J

| Code | Status | Source | Details |
|------|--------|--------|---------|
| e_http_status_200 | error | HTTP | HTTP response shall have StatusCode 200, but it is 404 Not Found |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| w_aits_pem_certificate_chain | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

### https://shaken.spectrum.com/4d65efdb8a1ca366e9576c8fda747fa4.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_aits_pem_certificate_chain | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |


Generated: 27/10/2022 at 22:42:50