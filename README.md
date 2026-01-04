### A GOLANG CONCURRENCY WORKFLOW INVOLVING THE USE OF GOROUTINES AND CHANNELS USING THE ONGOING AFCON 2025 FOOTBALL GROUP PAIRINGS AS CONTEXT, WHERE THE PAIRINGS OF THE TOURNAMENT ARE RENDERED VIA A CONSUMER CHANNEL IN ASCENDING AND DESCENDING ORDER VIA THE USE OF API ENDPOINTS.

<img width="1253" height="863" alt="Image" src="https://github.com/user-attachments/assets/6db8711d-7287-46b9-a0e3-3bcf0ec90db8" />

STEP 1:

At this stage, the Producer, collates all the AFCON Group pairing API Endpoints and passes them to the Job Channel for further processing.


STEP 2:

The Job Channel Accepts the API Endpoints from the Producer 



STEP 3:

The Job Channel passes each Endpoint to different Workers, where the API Endpoints are queried for their values and each is concurrently processed and passed to a results channel



STEP 4:

The results channel accepts the concurrent results. At this stage, the workers are done making GET requests to each of the AFCON group pairings.



STEP 5:

At this stage, The Consumer, receives the values from the results channel and then processes them in to both Ascending order and Descending order based on the API GET request Endpoint call.



STEP 6:

This is a snapshot of the Descending order of the Group pairings after an API call has been made.



STEP 7:

This is a snapshot of the Ascending order of the Group pairings after an API call has been made.
