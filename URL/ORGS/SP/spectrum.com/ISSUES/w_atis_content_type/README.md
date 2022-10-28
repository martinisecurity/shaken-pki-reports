# STIR/SHAKEN Certificate Repository Compliance

## spectrum.com

Code: w_atis_content_type\
Source: ATIS-1000080\
Description: ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain

### https://shaken.spectrum.com/4d65efdb8a1ca366e9576c8fda747fa4.pem

| Code | Status | Source | Details |
|------|--------|--------|---------|
| w_atis_content_type | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| e_atis_cache_header | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| w_atis_pem_certificate_chain | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |


Generated: 28/10/2022 at 18:54:21