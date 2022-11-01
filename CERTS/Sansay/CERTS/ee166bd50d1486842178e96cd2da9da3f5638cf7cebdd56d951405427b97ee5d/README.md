# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Terra Nova Telecom 382G

Tested At: 01 Nov 22 20:26 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=SHAKEN Terra Nova Telecom 382G, OU=NOC, O=Terra Nova Telecom, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/382G_TERRANOVA_STIR_SHAKEN.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDhjCCAyugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3QwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjIyNTMzMloXDTIyMTEyNTIyNTMzMlowczELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGzAZBgNVBAoMElRlcnJhIE5vdmEgVGVsZWNvbTEMMAoGA1UECwwDTk9DMScwJQYDVQQDDB5TSEFLRU4gVGVycmEgTm92YSBUZWxlY29tIDM4MkcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQUZcBpTXzq%2F4l2CuZ90OBlLvIaqacu478cFveXc0%2Bt6sPgddYepZCyT9JcZG0NYLjUHUjCQ9WZslg8AbwbS0Qeo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDM4MkcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQZ3CGz9usyXL1owkd%2FQ8Oe1gPMCjCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQC46hQH7%2BvhJAhIcGmPw2YnjJlj2eNQ36OukTpPCW%2Bo7QIhAIHlGuXewti84PDzV%2BLRJTqO5tMbE3HZydDQxz6zWdTR)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 382G' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 01/11/2022 at 20:34:21