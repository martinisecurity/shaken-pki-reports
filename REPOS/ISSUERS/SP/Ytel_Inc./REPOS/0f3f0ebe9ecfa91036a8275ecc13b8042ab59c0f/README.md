# STIR/SHAKEN Certificate Repository Compliance

## Ytel Inc.

Name: `https://s3.amazonaws.com/certificates.peeringhub.io/ytel/703J.crt`\
Tested At: 02 Nov 22 20:01 UTC\
Time: 16ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 02 Nov 22 20:09 UTC