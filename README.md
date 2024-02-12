# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 10224 certificates were included in the corpus being tested
- 824 certificates in the corpus were skipped because they are duplicates
- 8048 certificates in the corpus were skipped because they are expired
- 474 certificates in the corpus were skipped because they are not currently trusted
- 878 certificates being tested against the remaining rules
- 1.75 issues on average found in unexpired, trusted, and non-compliant certificates
- 64.46% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 8.77% of certificates are too old to be assessed against currently enforced expectations
- 271 days is the average remaining validity for the certificates in the corpus
- 271 days is the average initial validity for the certificates in the corpus
- 301 certificates expire in the next 30 days
- 14.86 average number of unexpired certificates per OCN observed
- 688 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 46 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 7 certificates in the corpus were skipped because they are not currently trusted
- 39 certificates being tested against the remaining rules
- 2.38 issues on average found in unexpired, trusted, and non-compliant certificates
- 41.03% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 74.36% of certificates are too old to be assessed against currently enforced expectations
- 5775 days is the average remaining validity for the certificates in the corpus
- 5694 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 56.95% of certificate repository URLs contain one or more Error level issue
- 58.54% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 25 (2.85%) | 25 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 4 (0.46%) | 4 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 308 (35.08%) | 11 (3.57%) | 0 (0.00%) | 0 (0.00%) | 1 (0.32%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 74 (8.43%) | 74 (100.00%) | 0 (0.00%) | 0 (0.00%) | 47 (63.51%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 6 (0.68%) | 6 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 157 (17.88%) | 157 (100.00%) | 0 (0.00%) | 0 (0.00%) | 26 (16.56%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 20 (2.28%) | 11 (55.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 24 (2.73%) | 24 (100.00%) | 0 (0.00%) | 0 (0.00%) | 1 (4.17%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 205 (23.35%) | 205 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (0.98%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.11%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Telonium](CERTS/Telonium/README.md#leaf-certificates) | 21 (2.39%) | 15 (71.43%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 33 (3.76%) | 33 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| **Total** | 878 (100.00%) | 566 (64.46%) | 0 (0.00%) | 0 (0.00%) | 77 (8.77%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [CTIA](CERTS/CTIA/README.md#ca-certificates) | 1 (2.56%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (5.13%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 3 (7.69%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 4 (10.26%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (50.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 3 (7.69%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (7.69%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (15.38%) | 1 (16.67%) | 0 (0.00%) | 0 (0.00%) | 6 (100.00%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (5.13%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (5.13%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (5.13%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 4 (10.26%) | 1 (25.00%) | 0 (0.00%) | 0 (0.00%) | 4 (100.00%) |
| [Telonium](CERTS/Telonium/README.md#ca-certificates) | 5 (12.82%) | 3 (60.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 2 (5.13%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| **Total** | 39 (100.00%) | 16 (41.03%) | 0 (0.00%) | 0 (0.00%) | 29 (74.36%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 12 Feb 24 17:02 UTC