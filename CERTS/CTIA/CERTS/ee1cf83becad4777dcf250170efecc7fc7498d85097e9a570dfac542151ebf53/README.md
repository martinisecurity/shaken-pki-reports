# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Root CA

Tested At: 21 Nov 23 01:54 UTC\
Initial Validity Period: 9125 day(s)\
Remaining Validity Period: 8973 day(s)\
Subject: CN=SHAKEN Root CA, O=CTIA, C=US\
Issuer: CN=SHAKEN Root CA, O=CTIA, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIB7DCCAXGgAwIBAgIUC121W1K%2BmtMXtilUY%2FzZ5sjGiD0wCgYIKoZIzj0EAwMwNTELMAkGA1UEBhMCVVMxDTALBgNVBAoMBENUSUExFzAVBgNVBAMMDlNIQUtFTiBSb290IENBMB4XDTIzMDYyMTEzMTUyNloXDTQ4MDYxNDEzMTUyNlowNTELMAkGA1UEBhMCVVMxDTALBgNVBAoMBENUSUExFzAVBgNVBAMMDlNIQUtFTiBSb290IENBMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEmY0HRLtHuM3vLsScB%2BJvKxyPJI420XSU8yrm2uJGt0q5c1HGJmntrC%2FdSC3tlbNTwm9UJAd41tZwvvwW6SUS4gHP%2FvZrdQY%2BXuofInFBt3rfmJNHih%2B16gCe9PpHN5Uvo0IwQDAPBgNVHRMBAf8EBTADAQH%2FMB0GA1UdDgQWBBTK97ldmVZoKMTO8hrjALPsmHwCJjAOBgNVHQ8BAf8EBAMCAgQwCgYIKoZIzj0EAwMDaQAwZgIxAOKB%2BmHi2xQKTN0%2Fa%2BPA6idgeZG5kPLjGrx83JDultsGizrVsEx0eje7y0eyKDRPFQIxAIEL8D9gfcMAD%2FD5sWQxWaO7kmofI%2FbVDwf6Wq3QkyLFtJ224bk%2BJ9qWuTGoOFBg9g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_root](../../ISSUES/e_atis_subject_cn_root/README.md) | error | ATIS1000080 | Common Name attribute 'SHAKEN Root CA' does not include the text string 'ROOT' (case insensitive). |
| [e_atis_ca_signature_algorithm](../../ISSUES/e_atis_ca_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.10045.4.3.3 |
| [e_atis_ca_subject_public_key](../../ISSUES/e_atis_ca_subject_public_key/README.md) | error | ATIS1000080 | Subject Public Key Info field contains a public key that is not 256 bits |


Generated: 21 Nov 23 01:55 UTC