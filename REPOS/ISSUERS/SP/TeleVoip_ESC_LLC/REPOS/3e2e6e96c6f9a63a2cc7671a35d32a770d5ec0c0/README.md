# STIR/SHAKEN Certificate Repository Compliance

## TeleVoip ESC LLC

Name: `https://ssc.getsipnav.com/certs/f5b45519d55a49e91c66b17a24bca2c98ee8d867`\
Tested At: 04 Oct 24 16:23 UTC\
Time: 119ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC