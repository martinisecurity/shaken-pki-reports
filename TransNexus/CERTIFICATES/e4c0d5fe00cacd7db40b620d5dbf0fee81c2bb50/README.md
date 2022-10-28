# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate e4c0d5fe00cacd7db40b620d5dbf0fee81c2bb50
Tested At: 2022-10-28 18:54:56 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 6 day(s)\
Subject: CN=SHAKEN 459J, OU=SHAKEN, O=Altaworx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/d3d6dbfc-2914-49c7-8f47-d0aa5bd5d764/6e0df5aab72d070417858423b3f50a35.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6jCCApCgAwIBAgIQfaJsozqcqIALiQ0b5bOdrjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDIzNTZaFw0yMjExMDIyMDIzNTVaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhBbHRhd29yeDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNDU5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDLfF9e%2B3SK31blHxOO2SmqwGb%2F2WZe%2FhL3S%2FKLqTduKmI9F8aQ4T8IO1dDCX4RGCN5yce9OJu%2Fop%2FbrjAyNK2KjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUNlZlVgr99IjoEh7SnHb0OQegxWswHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENDU5SjAKBggqhkjOPQQDAgNIADBFAiEAtUs79gicJ43zqGSmYqtPalf4YMRJM65Us%2FxDekKnXU0CIEBLMUjRAKl9V8fTGwjD6ulZt8b%2BmOs5vO8qRd5hQlES)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 18:55:01