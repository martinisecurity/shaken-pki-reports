# STIR/SHAKEN Certificate Repository Compliance

## Voip Stir PR LLC

Name: `https://ssc.getsipnav.com/certs/ffad5de97793861e660be844cf52d740aa256d7d`\
Tested At: 04 Oct 24 16:23 UTC\
Time: 131ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC