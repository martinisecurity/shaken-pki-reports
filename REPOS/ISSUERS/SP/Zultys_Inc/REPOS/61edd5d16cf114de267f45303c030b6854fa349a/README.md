# STIR/SHAKEN Certificate Repository Compliance

## Zultys Inc

Name: `https://zultys-pem-cert-2022.s3.amazonaws.com/77c9a883d0add03d43f2534413f32cba.pem`\
Tested At: 27 Nov 23 22:52 UTC\
Time: 334ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 27 Nov 23 22:56 UTC