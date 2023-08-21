# STIR/SHAKEN Certificate Repository Compliance

## Unknown

Name: `http://5.161.195.139/ec256-public.pem`\
Tested At: 21 Aug 23 19:59 UTC\
Time: 421ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [e_atis_redirect](../../ISSUES/e_atis_redirect/README.md) | error | ATIS-1000074 | The STI-VS shall not follow HTTP redirections |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_pem_certificate_chain](../../ISSUES/w_atis_pem_certificate_chain/README.md) | warn | ATIS-1000080 | HTTP response body should be PEM certificate chain. Response body is not PEM encoded |

Generated: 21 Aug 23 20:18 UTC