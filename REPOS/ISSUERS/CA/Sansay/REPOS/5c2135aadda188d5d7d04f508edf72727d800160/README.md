# STIR/SHAKEN Certificate Repository Compliance

## Sansay

Name: `https://cr.sansay.com/363A/order/142_363A_17`\
Tested At: 12 Feb 24 19:03 UTC\
Time: 278ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 19:26 UTC