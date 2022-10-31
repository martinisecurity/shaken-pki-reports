# STIR/SHAKEN Certificate Repository Compliance

## Zultys Inc

Name: `https://zultys-pem-cert-2022.s3.amazonaws.com/89179fa533bbaf0aea20a9f31aa06f20.pem`\
Tested At: 31 Oct 22 19:21 UTC\
Time: 324ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 31/10/2022 at 19:21:49