# STIR/SHAKEN Certificate Repository Compliance

## MYPBXManager LLC

Name: `https://www.mypbxmanager.net/local/PubShakenCert.crt`\
Tested At: 23 Nov 22 18:09 UTC\
Time: 73ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 23 Nov 22 18:09 UTC