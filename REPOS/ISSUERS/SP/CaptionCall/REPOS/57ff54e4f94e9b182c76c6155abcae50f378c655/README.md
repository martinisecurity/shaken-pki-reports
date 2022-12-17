# STIR/SHAKEN Certificate Repository Compliance

## CaptionCall

Name: `https://api.alianza.com/v2/stir-shaken/certs/2dd72c30-5e06-49a5-bbec-fff3cdf9444b/key.crt`\
Tested At: 17 Dec 22 16:57 UTC\
Time: 357ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 17 Dec 22 17:07 UTC