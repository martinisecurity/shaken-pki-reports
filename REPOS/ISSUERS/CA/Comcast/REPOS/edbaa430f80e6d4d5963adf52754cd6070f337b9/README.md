# STIR/SHAKEN Certificate Repository Compliance

## Comcast

Name: `https://sticr.stir.comcast.com/2a3f24ff-d20d-491c-ae67-4d667204fbb7.cer`\
Tested At: 04 Oct 24 16:25 UTC\
Time: 35ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_http_status_200](../../ISSUES/e_http_status_200/README.md) | error | HTTP | HTTP response shall have StatusCode 200, but it is 404 Not Found |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 04 Oct 24 16:29 UTC