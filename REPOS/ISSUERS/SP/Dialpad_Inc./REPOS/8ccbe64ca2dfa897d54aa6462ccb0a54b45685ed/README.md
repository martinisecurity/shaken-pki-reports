# STIR/SHAKEN Certificate Repository Compliance

## Dialpad Inc.

Name: `https://fstelephony.appspot.com/static/cert/shaken_identity_2022_3.cer`\
Tested At: 31 Oct 22 18:33 UTC\
Time: 91ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 31/10/2022 at 18:34:12