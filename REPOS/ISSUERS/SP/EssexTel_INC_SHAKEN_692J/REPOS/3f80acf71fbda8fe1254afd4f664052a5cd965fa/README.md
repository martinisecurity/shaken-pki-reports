# STIR/SHAKEN Certificate Repository Compliance

## EssexTel INC SHAKEN 692J

Name: `https://netnumber-sti-cr.s3.amazonaws.com/certs/041091ff-fdf3-400e-b212-fe7f448439bd`\
Tested At: 06 Jul 23 14:06 UTC\
Time: 75ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 06 Jul 23 14:08 UTC