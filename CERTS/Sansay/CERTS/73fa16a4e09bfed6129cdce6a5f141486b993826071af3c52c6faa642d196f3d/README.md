# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Matrix 9451

Tested At: 31 Oct 22 20:40 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 26 day(s)\
Subject: CN=SHAKEN Matrix 9451, OU=Engineering, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Lingo-9451

View: [Click to view](https://understandingwebpki.com/?cert=MIICxjCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU4kwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwNDI1N1oXDTIyMTEyNjAwNDI1N1owYTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxFDASBgNVBAsMC0VuZ2luZWVyaW5nMRswGQYDVQQDDBJTSEFLRU4gTWF0cml4IDk0NTEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATKK19BR8TbDyCPk5h2YzHgsh%2BJs%2BPyP%2FPJ%2F4T5HwXs7u43XyL9Z%2BZZ2RNErU5qraZmexjDBbnguUi0kNUn%2B59So4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5NDUxMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUK5Ds9QZFxG%2F%2BcfvTtwaOpOm5qd0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCZqvq%2BP%2BrtMjziu4OkvQ8%2FbyxSUTT106NFW%2B7fO%2B6agwIhALfmmHL%2F2jlXuopl5sTpEIPm%2BhkwqNGhpNj7aT4ThYLZ)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 9451' |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 20:47:45