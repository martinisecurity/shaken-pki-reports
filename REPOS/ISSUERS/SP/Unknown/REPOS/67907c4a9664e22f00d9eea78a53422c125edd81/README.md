# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Name: `http://5.161.195.139/ec256-public.pem`\
Tested At: 15 Nov 23 16:04 UTC\
Time: 449ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [e_atis_redirect](../../ISSUES/e_atis_redirect/README.md) | error | ATIS-1000074 | The STI-VS shall not follow HTTP redirections |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 15 Nov 23 17:17 UTC