# STIR/SHAKEN CA Ecosystem Compliance

## Certificate HD CARRIER LLC

Tested At: 12 Apr 23 21:51 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 20 day(s)\
Subject: L=Long Beach, ST=California, O=HD CARRIER LLC, C=US, CN=HD CARRIER LLC\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/1eff5935-82b7-4df5-a56d-2f870e2b1ee8

[View certificate details](https://understandingwebpki.com/?cert=MIICpTCCAiugAwIBAgIJALjTeVUvsf51MAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzA0MDIwMDAwMDBaFw0yMzA1MDEyMzU5NTlaMGkxFzAVBgNVBAMMDkhEIENBUlJJRVIgTExDMQswCQYDVQQGEwJVUzEXMBUGA1UECgwOSEQgQ0FSUklFUiBMTEMxEzARBgNVBAgMCkNhbGlmb3JuaWExEzARBgNVBAcMCkxvbmcgQmVhY2gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2BjoiWl6irYziXFAhlLCYvdJM89omNmggSAFtV%2F6DLW2LIqknZPy7jK3ki%2Fj1x9MSrlQ71x66y2t16Da5xroQto4GVMIGSMBYGCCsGAQUFBwEaBAowCKAGFgQzMjFKMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBRO4uQ%2Ffm72A4oT%2Fr6226RSEv8OhTAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIxANVFCVSQrjmhU25VOfGsyy2g2Hm6x8%2Br%2BRwmFB7GEBl4S1pXfX%2FwSa8zmcq91QVi8QIwK6E7ZAM%2BQoU5kFnhZOrTBQWf%2BYfoRjtxggrS7RMtpmZRkDRnAY7OSN8M4LsT92bn)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 321J' |


Generated: 12 Apr 23 22:02 UTC