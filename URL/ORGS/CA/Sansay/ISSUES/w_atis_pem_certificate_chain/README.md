# STIR/SHAKEN Certificate Repository Compliance

## Sansay

Code: w_atis_pem_certificate_chain\
Source: ATIS-1000080\
Description: ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain

### https://cr.sansay.com/Mitel_Cloud_Services_Inc._670J

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_http_status_200 | error | HTTP | HTTP response shall have StatusCode 200, but it is 404 Not Found |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_pem_certificate_chain | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |


Generated: 28/10/2022 at 18:22:16