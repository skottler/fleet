# Sales

The Fleet sales team embodies [our values](https://fleetdm.com/handbook/company#values) in every aspect of our work. Specifically, we continuously work to overperform and achieve strong results. We prioritize efficiency in our processes and operations. We succeed because of transparent, cross-functional collaboration. We are committed to hiring for and celebrating diversity, and we strive to create an environment of inclusiveness and belonging for all. We embrace a spirit of iteration, understanding that we can always improve.

## Outreach one-pager

Our one-pager offers a summary of what Fleet does. It can help stakeholders become familiar with the company and product while also being a useful tool the Growth team uses for sales outreach. Find Fleet's outreach one-pager in this [Google Doc](https://docs.google.com/presentation/d/1GzSjUZj1RrRBpa_yHJjOrvOTsldQQKfq927vpKP1lpU/edit?usp=sharing).

## Intro deck

Fleet's intro deck adds additional detail to our pitch. Find it in [Google Slides](https://docs.google.com/presentation/d/1GzSjUZj1RrRBpa_yHJjOrvOTsldQQKfq927vpKP1lpU/edit?usp=sharing).

## Intro video

Fleet's intro video shows how to get started with Fleet as an admin. Find it on [YouTube](https://www.youtube.com/watch?v=rVxSgvKjrWo).

## SOC 2

You can find a copy of Fleet's SOC 2 report in [Google Drive](https://drive.google.com/file/d/1B-Xb4ZVmZk7Fk0IA1eCr8tCVJ-cfipid/view?usp=drivesdk).  In its current form, this SOC 2 report is intended to be shared only with parties who have signed a non-disclosure agreement with Fleet.

You can learn more about how Fleet approaches security in the [security handbook](https://fleetdm.com/handbook/security) or in [Fleet's trust report](https://fleetdm.com/trust).

## Our lead handling and outreach approach

Fleet's main source for prospects to learn about the company and its offerings is our website, fleetdm.com. There are many places across the website for prospects to ask for more information, request merchandise, try the product and even purchase licenses directly. If the user experience in any of these locations asks for an email address or other contact information, Fleet may use that contact information for follow-up, including sales and marketing purposes. That contact information is for Fleet's sole use, and we do not give or sell that information to any third parties.

In the case of a prospect or customer request, we strive to adhere to the following response times:
- Web chat: 1 hour response during working hours, 8 hours otherwise
- Talk to an expert: prospects can schedule chats via our calendar tool
- All other enquiries: 1-2 days

Fleet employees can find other expectations for action and response times in this [internal document](https://docs.google.com/presentation/d/104-TRXlY55g303q2xazY1bpcDx4dHqS5O5VdJ05OwzE/edit?usp=sharing)

## Salesforce lead status flow

To track the stage of the sales cycle that a lead is at, we use the following standardized lead statuses to indicate which stage of the sales process a lead is at.
|Lead status                 | Description                                         |
|:-----------------------------|:----------------------------------------------------|
| New | Default status for all new leads when initially entered into Salesforce. We have an email or LinkedIn profile URL for the lead, but no established intent. The lead is just a relevant person to reach out to.|
| New enriched | Fleet enriched the lead with additional contact info.|
| New MQL | Lead has been established as a marketing qualified lead, meeting company size criteria.|
| Working to engage | Fleet (often Sales development representative-SDR) is working to engage the lead. |
| Engaged | Fleet has successfully made contact with the lead |
| Meeting scheduled | Fleet has scheduled a meeting with the lead. |
| Working to convert | Not enough info on Lead's Budget, Authority, Need and Timing (BANT) to be converted into an opportunity. |
| Closed nurture | Lead does not meet BANT criteria to be converted to an opportunity, but we should maintain contact with the lead as it may be fruitful in the future. |
| Closed do not contact | Lead does not meet BANT criteria for conversion, and we should not reach out to them again. |
| SAO Converted | Lead has met BANT criteria and successfully converted to an opportunity. |

At times, our sales team will reach out to prospective customers before they come to Fleet for information. Our cold approach is inspired by Daniel Grzelak’s (Founder, investor, advisor, hacker, CISO) [LinkedIn post](https://www.linkedin.com/posts/danielgrzelak_if-you-are-going-to-do-a-cold-approach-be-activity-6940518616459022336-iYE7). The following are the keys to an engaging cold approach. Since cold approaches like these can be easily ignored as mass emails, it’s important to personalize each one. 

- Research each prospect.
- Praise what’s great about their company.
- Avoid just stating facts about our product.
- State why we would love to work with them.
- Ask questions about their company and current device management experience.
- Keep an enthusiastic and warm tone.
- Be personable.
- Ask for the meeting with a proposed time.

Importantly, when we interact with CISOs or, for that matter, any member of a prospective customer organization, we adhere to the principles in this [LinkedIn post](https://www.linkedin.com/pulse/selling-ciso-james-turner). Specifically:

- Be curteous
- Be honest
- Show respect
- Build trust
- Grow relationships
- Help people

## Sales team writing principles

When writing for the Sales team, we want to abide by the following principles in our communications.

### Maintain naming conventions

Maintain naming conventions so people can expect what fields will look like when revisiting automations outside of Salesforce. This helps them avoid misunderstanding jargon and making mistakes that break automated integrations and cause business problems. One way we do this is by using sentence case where only the first word is capitalized (unless it’s a proper noun). See the below examples.

| Good job! ✅          | Don't do this. ❌    |
|:----------------------|:---------------------|
| Bad data              | Bad Data

### Be explicit

Being explicit helps people to understand what they are reading and how to use terms for proper use of automations outside of Salesforce. In the case of acronyms, that means expanding and treating them as proper nouns. Note the template for including acronyms is in the first column below.

| Good job! ✅          | Don't do this. ❌    |
|:----------------------|:---------------------|
| Do Not Contact (DNC)  | DNC



## Salesforce contributor experience checkups

In order to maintain a consistent contributor experience in Salesforce, we log in to make sure the structure of Salesforce data continues to look correct based on processes started elsewhere. Then we can look and see that the goals we want to achieve as a business are in line with our view inside Salesforce by conducting the following checkup. Any discrepancies between how information is presented in Salesforce and what should be in there per this ritual should be flagged so that they can be fixed or discussed.

1. Make sure the default tabs for a standard user include a detailed view of contacts, opportunities, accounts, and leads. No other tabs should exist.

2. Click the accounts tab and check for the following: 

* The default filter is Customers when you click on the accounts tab. Click on an account to continue.
* Click on a customer and make sure billing address, parent account, LinkedIn company URL, CISO employees (#), employees, and industry appear first at the top of the account.
* "Looking for meeting notes" reminder should appear on the right of the screen.  
* Useful links section should include links to Purchase Orders (POs), signed subscription agreements, invoices sent, meeting notes, and signed NDA. Clicking these links should search the appropriate repository for the requested information pertaining to the customer.
* Additional information section should include fields for account (customer) name first, account rating, LinkedIn sales navigator URL, LinkedIn company URL, and my LinkedIn overlaps. Make sure the LinkedIn links work.
* Accounting section should include the following fields: invoice sent (latest), the payment received on (latest), subscription end date (latest), press approval field, license key, total opportunities (#), deals won (#), close date (first deal), cumulative revenue, payment terms, billing address, and shipping address. 
* Opportunities, meeting notes, and activity feed should appear on the right.  

3. Click on the opportunities tab and check for the following:

* Default filter should be all opportunities. Open an opportunity to continue.
* Section at the top of the page should include fields for account name, amount, close date, next step, and opportunity owner.
* Opportunity information section should include fields for account name, opportunity name (should have the year on it), amount, next step, next step's due date, close date, and stage.
* The accounting section here should include: up to # of hosts, type, payment terms, billing process, term, reseller, effective date, subscription end date, invoice sent, and the date payment was received.
* Stage history, activity feed, and LinkedIn sales navigator should appear at the right.  

4. Click on the contacts tab and check for the following:

* Default filter should be all contacts. Open a contact to continue.
* Top section should have fields for the contact's name, job title, department, account name, LinkedIn, and Orbit feed. 
* The second section should have fields for LinkedIn URL, account name, name, title, is champion, and reports to
* Additional information should have fields for email, personal email, Twitter, GitHub, mobile, website, orbit feed, and description.
* Related contacts section should exist at the bottom, activity feed, meeting notes reminder, and manager information should appear on the right. 

5. Click on the leads tab and check for the following:

* Default filter should be all leads. Open a lead to continue.
* There should be fields for name, lead source, lead status, and rating.

## Rituals

Directly Responsible Individuals (DRI) engage in the ritual(s) below at the frequency specified.

| Ritual                       | Frequency                | Description                                         | DRI               |
|:-----------------------------|:-----------------------------|:----------------------------------------------------|-------------------|
| Sales huddle | Weekly | Agenda: Go through every [open opportunity](https://fleetdm.lightning.force.com/lightning/o/Opportunity/list?filterName=00B4x00000CTHZIEA5) and update the next steps. | Alex Mitchell
[Salesforce contributor experience checkup](#salesforce-contributor-experience-checkups)| Monthly | Make sure all users see a detailed view of contacts, opportunities, accounts, and leads. | Nathan Holliday |
| Lead pipeline review  | Weekly | Agenda: Review leads by status/stage; make sure SLAs are met. | Alex Mitchell 
| TODO  | TODO | TODO | TODO 

## Slack channels

This group maintains the following [Slack channels](https://fleetdm.com/handbook/company#why-group-slack-channels):

| Slack channel                       | [DRI](https://fleetdm.com/handbook/company#why-group-slack-channels)    |
|:------------------------------------|:--------------------------------------------------------------------|
| `#g-sales`                     | Alex Mitchell
| `#_from-prospective-customers` | Alex Mitchell




<meta name="maintainedBy" value="alexmitchelliii">
<meta name="title" value="🤝 Sales">
