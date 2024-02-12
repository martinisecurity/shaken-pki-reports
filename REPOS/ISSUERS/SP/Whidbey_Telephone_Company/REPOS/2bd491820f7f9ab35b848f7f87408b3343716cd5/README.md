# STIR/SHAKEN Certificate Repository Compliance

## Whidbey Telephone Company

Name: `https://cdn-cr.cgah.tnsi.com/certs/4b8c18a5a68450eb480b9727e86382fdaa4afa66`\
Tested At: 12 Feb 24 16:26 UTC\
Time: 10ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC