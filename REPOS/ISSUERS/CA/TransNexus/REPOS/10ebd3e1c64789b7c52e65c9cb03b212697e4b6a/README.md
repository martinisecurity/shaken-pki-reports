# STIR/SHAKEN Certificate Repository Compliance

## TransNexus

Name: `https://certificates.clearip.com/f8a68e44-8fb0-4f28-b533-c4df27ed8e1b/7c64b4c30b55c0c3265b7503afeb1b89.pem`\
Tested At: 04 Oct 24 15:47 UTC\
Time: 83ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_http_status_200](../../ISSUES/e_http_status_200/README.md) | error | HTTP | HTTP response shall have StatusCode 200, but it is 503 Service Unavailable |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 04 Oct 24 16:29 UTC