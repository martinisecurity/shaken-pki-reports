# STIR/SHAKEN Certificate Repository Compliance

## Apex Telecom

Name: `https://s3.amazonaws.com/certificates.peeringhub.io/288K/288K.crt`\
Tested At: 31 Oct 22 16:42 UTC\
Time: 179ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 31/10/2022 at 16:43:22