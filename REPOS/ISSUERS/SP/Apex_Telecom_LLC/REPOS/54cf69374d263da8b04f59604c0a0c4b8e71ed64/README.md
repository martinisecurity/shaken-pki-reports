# STIR/SHAKEN Certificate Repository Compliance

## Apex Telecom LLC

Name: `https://s3.amazonaws.com/certificates.peeringhub.io/288K/288K.crt`\
Tested At: 02 Nov 22 17:24 UTC\
Time: 43ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 02 Nov 22 17:25 UTC