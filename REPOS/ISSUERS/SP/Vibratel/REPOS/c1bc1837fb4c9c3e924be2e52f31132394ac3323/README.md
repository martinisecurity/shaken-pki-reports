# STIR/SHAKEN Certificate Repository Compliance

## Vibratel

Name: `https://ssc.getsipnav.com/certs/2c80ab6c6b60bccb742a11f1ac57540f8ad29097`\
Tested At: 21 Aug 23 20:16 UTC\
Time: 56ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Aug 23 20:18 UTC