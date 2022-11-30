# STIR/SHAKEN Certificate Repository Compliance

## Metaswitch

Name: `https://sti-cr.cgah.tnsi.com/certs/6c54d88f5669c28848983d4572f11c0b896342ac`\
Tested At: 30 Nov 22 17:30 UTC\
Time: 15204ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 30 Nov 22 17:39 UTC