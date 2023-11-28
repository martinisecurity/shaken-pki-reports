# STIR/SHAKEN Certificate Repository Compliance

## West Kentucky Rural Telephone Cooperative

Name: `https://cdn-cr.cgah.tnsi.com/certs/fecae0c1f85acf27b0d6f8afbe4c90593fd6ecc7`\
Tested At: 28 Nov 23 15:59 UTC\
Time: 21ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 23 16:15 UTC