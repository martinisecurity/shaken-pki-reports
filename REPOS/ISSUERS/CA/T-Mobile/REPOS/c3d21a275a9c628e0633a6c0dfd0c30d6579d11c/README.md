# STIR/SHAKEN Certificate Repository Compliance

## T-Mobile

Name: `https://t-mobile-sticr.fosrvt.com/d1ed2fcf74511801e3df6deffbf762b764221487a8305a7da32e9efab2cbf358.pem`\
Tested At: 15 Nov 23 18:09 UTC\
Time: 27ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 15 Nov 23 18:10 UTC