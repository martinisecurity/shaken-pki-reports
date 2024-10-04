# STIR/SHAKEN Certificate Repository Compliance

## Zultys Inc

Name: `https://zultys-pem-cert-2022.s3.amazonaws.com/91ac6d5d106cc23ac5aba3b88594c59f.pem`\
Tested At: 04 Oct 24 16:26 UTC\
Time: 87ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC