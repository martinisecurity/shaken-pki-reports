# STIR/SHAKEN Certificate Repository Compliance

## TeleVoip ESC LLC

Name: `https://ssc.getsipnav.com/certs/f5b45519d55a49e91c66b17a24bca2c98ee8d867`\
Tested At: 24 Nov 23 11:11 UTC\
Time: 145ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 24 Nov 23 11:17 UTC