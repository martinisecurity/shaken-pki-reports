# STIR/SHAKEN Certificate Repository Compliance

## T-Mobile

Name: `https://t-mobile-sticr.fosrvt.com/4816be075b3ad7d142c4cdb94bba30508bbb8022d32962ac8ab87e454333020f.pem`\
Tested At: 18 Aug 25 21:09 UTC\
Time: 134ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC