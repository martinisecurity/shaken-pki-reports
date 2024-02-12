# STIR/SHAKEN Certificate Repository Compliance

## ATT SHAKEN E-E

- 3 repository URLs were included in the corpus being tested
- 1 repository URLs in the corpus were skipped because they were duplicated
- 2 repository URLs being tested against the remaining rules
- 2.00 issues on average found in non-compliant certificate repository URLs
- 100.00% of repository URLs contain one or more Error level issue
- 100.00% of repository URLs contain one or more Warning level issue
- 0.00% of repository URLs contain one or more Notice level issue
- 221ms average time it took to download each certificate

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_cache_header](ISSUES/e_atis_cache_header/README.md) | ATIS-1000074 |
| 2 | [w_atis_content_type](ISSUES/w_atis_content_type/README.md) | ATIS-1000080 |

| Repository URLs | Not After |  Problems | Link |
|-----------------|-----------|-----------|------|
| `https://cert.sticr.att.net:8443/certs/att/bf0e7932-fe6e-4a5a-9948-adafb1330487` | 06&#160;Dec&#160;24&#160;16:28&#160;UTC | true | [view](REPOS/a25ba5122850a463cd859dceaa16b5d061abc0ba/README.md) |
| `https://cert2.sticr.att.net:8443/sti-cr/att-stica1691767810461-cert.crt` | 10&#160;Aug&#160;24&#160;15:15&#160;UTC | true | [view](REPOS/e927f896f26103d706a2a5a669617a09cb847758/README.md) |


Generated: 12 Feb 24 19:26 UTC