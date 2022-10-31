# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 157C

Tested At: 31 Oct 22 18:32 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 157C, OU=SHAKEN, O=DigitalSpeed Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/82d97ecb-9b38-4926-9062-8e21b8986930/4a5ab5a3b56a2bb35ff2b29c212a2ab8.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FTCCAqOgAwIBAgIQUc%2FTIDaxT%2F2Jw1yE1UykHjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjQyMDUxNDdaFw0yMjEwMzEyMDUxNDZaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtEaWdpdGFsU3BlZWQgQ29tbXVuaWNhdGlvbnMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE1N0MwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQtQtb5sgxAvrvZkWelm1wCwBrPmgw%2FAAKpLWW9rvdcQ1IaC%2B9s6jGF%2Fu399EDAixu%2BDrRE7GxN9N1E5mWvQMkco4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFCeG6sRn%2F%2F3KU%2BjOkGVue5T3Lys6MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE1N0MwCgYIKoZIzj0EAwIDSAAwRQIgMg5zgnvjZOm6AOGf9b%2FRhECe%2Bs1g5UAIQpa6p%2Br3T5oCIQCSGvVyToWMomZcZ1L6jsL1eChZ2MupKh1Sa2Qss13P2Q%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31/10/2022 at 18:34:12