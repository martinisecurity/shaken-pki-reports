# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 8876 certificates were included in the corpus being tested
- 734 certificates in the corpus were skipped because they are duplicates
- 6781 certificates in the corpus were skipped because they are expired
- 474 certificates in the corpus were skipped because they are not currently trusted
- 887 certificates being tested against the remaining rules
- 3.52 issues on average found in unexpired, trusted, and non-compliant certificates
- 57.05% of certificates contain one or more Error level issue
- 27.85% of certificates contain one or more Warning level issue
- 0.56% of certificates contain one or more Notice level issue
- 6.76% of certificates are too old to be assessed against currently enforced expectations
- 266 days is the average remaining validity for the certificates in the corpus
- 266 days is the average initial validity for the certificates in the corpus
- 300 certificates expire in the next 30 days
- 13.53 average number of unexpired certificates per OCN observed
- 656 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 45 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 8 certificates in the corpus were skipped because they are not currently trusted
- 37 certificates being tested against the remaining rules
- 2.08 issues on average found in unexpired, trusted, and non-compliant certificates
- 51.35% of certificates contain one or more Error level issue
- 37.84% of certificates contain one or more Warning level issue
- 2.70% of certificates contain one or more Notice level issue
- 54.05% of certificates are too old to be assessed against currently enforced expectations
- 5648 days is the average remaining validity for the certificates in the corpus
- 5558 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 35.06% of certificate repository URLs contain one or more Error level issue
- 54.79% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 18 (2.03%) | 18 (100.00%) | 18 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 2 (0.23%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 376 (42.39%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 71 (8.00%) | 71 (100.00%) | 1 (1.41%) | 0 (0.00%) | 41 (57.75%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 5 (0.56%) | 5 (100.00%) | 1 (20.00%) | 5 (100.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 148 (16.69%) | 148 (100.00%) | 22 (14.86%) | 0 (0.00%) | 19 (12.84%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 16 (1.80%) | 16 (100.00%) | 1 (6.25%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 25 (2.82%) | 25 (100.00%) | 25 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 178 (20.07%) | 178 (100.00%) | 178 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.11%) | 1 (100.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Telonium](CERTS/Telonium/README.md#leaf-certificates) | 15 (1.69%) | 15 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Telonium Communications LLC](CERTS/Telonium_Communications_LLC/README.md#leaf-certificates) | 5 (0.56%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 27 (3.04%) | 27 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| **Total** | 887 (100.00%) | 506 (57.05%) | 247 (27.85%) | 5 (0.56%) | 60 (6.76%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [CTIA](CERTS/CTIA/README.md#ca-certificates) | 1 (2.70%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (5.41%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (5.41%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (8.11%) | 0 (0.00%) | 0 (0.00%) | 1 (33.33%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 3 (8.11%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (8.11%) | 2 (66.67%) | 2 (66.67%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (16.22%) | 4 (66.67%) | 6 (100.00%) | 0 (0.00%) | 4 (66.67%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (5.41%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (5.41%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (5.41%) | 0 (0.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 4 (10.81%) | 1 (25.00%) | 0 (0.00%) | 0 (0.00%) | 3 (75.00%) |
| [Telonium](CERTS/Telonium/README.md#ca-certificates) | 3 (8.11%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Telonium Communications LLC](CERTS/Telonium_Communications_LLC/README.md#ca-certificates) | 2 (5.41%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 2 (5.41%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| **Total** | 37 (100.00%) | 19 (51.35%) | 14 (37.84%) | 1 (2.70%) | 20 (54.05%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 15 Nov 23 18:10 UTC