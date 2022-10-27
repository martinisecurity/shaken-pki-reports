# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 2661ebd23f777002d5a8b8e0cced45b06bf12b10
Tested At: 2022-10-27 00:05:33 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 291K, OU=SHAKEN, O=Hypercore Networks\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/b66a7496-cdd4-46e3-b219-cb0ecdc80d22/1925a83113629fb0818d64cdf29274e3.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BTCCAp%2BgAwIBAgIQd7hkpTJM5ILirAD4Zlf8tTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjQxNjMyMTdaFw0yMjEwMzExNjMyMTZaMFYxCzAJBgNVBAYTAlVTMSAwHgYDVQQKExdIeXBlcmNvcmUgTmV0d29ya3MsIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMjkxSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHXAySNyMGVlbKuhWfUpiMiadU5l5d%2F17Hjg%2Fpd5nlEel2lIDo35gXYRV17huUs1aexx%2FkDWWM1ibh9hSG2bCnSjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUd7Resfp3PXewOe1NTiE31vVtOF8wHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMjkxSzAKBggqhkjOPQQDAgNIADBFAiAsQEDUJS5Z3RRqDV5MSsZRt%2BwsnczaIUWWGrUjvgNO0QIhAIYGGl0Bar68aCoLHDjuRW6JB%2FxCkuQqBNZizixit1%2F5)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 00:07:07