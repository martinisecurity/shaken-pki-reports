# STIR/SHAKEN CA Ecosystem Compliance

## Certificate HD CARRIER LLC

Tested At: 15 Nov 23 16:11 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -153 day(s)\
Subject: L=Long Beach, ST=California, O=HD CARRIER LLC, C=US, CN=HD CARRIER LLC\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/092ead57-d2e2-47f4-8c47-98acdeefd503

[View certificate details](https://understandingwebpki.com/?cert=MIICpTCCAiugAwIBAgIJAM3Dfg7UCp0XMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzA1MTYwMDAwMDBaFw0yMzA2MTQyMzU5NTlaMGkxFzAVBgNVBAMMDkhEIENBUlJJRVIgTExDMQswCQYDVQQGEwJVUzEXMBUGA1UECgwOSEQgQ0FSUklFUiBMTEMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAcMCkxvbmcgQmVhY2gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2BjoiWl6irYziXFAhlLCYvdJM89omNmggSAFtV%2F6DLW2LIqknZPy7jK3ki%2Fj1x9MSrlQ71x66y2t16Da5xroQto4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQzMjFKMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBRO4uQ%2Ffm72A4oT%2Fr6226RSEv8OhTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIxAKxoxOwtpWUnZaVoOzuv4MJH6tgQANOAD1MmRizwwkh%2FFH0UlQNW%2BK4oXQNlni43GwIwPwg%2FYYkzvKDmpejULzvnLxPsoGkxNCYg8zA7V3AVgtVx6Bqdu%2Blji0cO0SYMwE0B)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 321J' |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 16:51 UTC