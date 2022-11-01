# STIR/SHAKEN Certificate Repository Compliance

## Whitesky Communications LLC

Name: `https://certs.sip.wtsky.net/prod/shaken-6-2023.cer`\
Tested At: 01 Nov 22 16:02 UTC\
Time: 145ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01/11/2022 at 16:11:23