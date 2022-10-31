# STIR/SHAKEN Certificate Repository Compliance

## Metaswitch

Name: `https://sti-cr.cgah.tnsi.com/certs/cd1856717574765eb6b4bddb7a5bc8814e1e2103`\
Tested At: 31 Oct 22 20:46 UTC\
Time: 15190ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 31/10/2022 at 20:47:45