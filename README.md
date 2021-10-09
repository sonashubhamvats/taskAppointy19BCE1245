# taskAppointy19BCE1245
## Sona Shubham Vats, 19BCE1245
I have created a HTTP JSON API using Go. The API has the following functionalities:-
<br>
* Create a user<br><br>
  The command for creating a user looks like this:-<br>
  <image src="https://user-images.githubusercontent.com/66525380/136666952-a4ded5de-c7a0-4736-9501-020689e9da03.PNG"><br><br>
  After sending the request the API creates a user by adding a field in the monogodb collection "users" with the information as     specified in the post command:-<br>
  <image src="https://user-images.githubusercontent.com/66525380/136667146-9e6a8625-ff17-41a2-9b51-f032b6afad34.PNG"><br>
  ![Capture4](https://user-images.githubusercontent.com/66525380/136667645-4686086c-b17d-46a2-bffd-3f30a1ce50b8.PNG)
  <br>
  The new user is created using the the command collection.InsertOne(ctx,user) in **CreateUserEnddPoint** function where user is     the decoded information taken from the post request and ctx is the context:-
 
* Get a user using user ID<br>
  The command for getting a user by its user_id looks like this, where the string after the user/ depicts the user_id that we need to   search for:-<br>
  ![Capture3](https://user-images.githubusercontent.com/66525380/136667471-7151f68a-f884-473f-8dd7-63e5e68412fc.PNG)
  <br><br>
  After sending the request the API first decodes the request to get the user_id that we need to search for using a regex and then     we search for the user with the specified user_id by iterating over the list of users using cursor, the output is then encoded into 
  the *http.ReponseWriter* , all of these functionalites are carried out in the function **GetUsersEndPoint**, the output looks like   this:-<br>
  ![Capture5](https://user-images.githubusercontent.com/66525380/136667696-53a575b8-3a35-4027-987f-8ab792dbca71.PNG)
* Create a post<br>
  Similar to create user , creating a post is carried out by a POST request and the functionality is carried out in the function     **CreatePostsEndpoint**, the command and output looks like this:-<br>
  ![Capture7](https://user-images.githubusercontent.com/66525380/136668004-4232f222-ac6a-430f-93cd-49085f7afc70.PNG)
  <br>
  ![Capture8](https://user-images.githubusercontent.com/66525380/136668012-21db2264-6d41-4acb-9da3-eb27dda05b03.PNG)
  <br>
  The database looks like this after execution:-<br>
  ![Capture6](https://user-images.githubusercontent.com/66525380/136668037-766945f9-8508-4dd7-b223-a030ea53b100.PNG)
* Get a post using post_id<br>
  Similar to getting a user by its user id we are retreiving the post_id from the url itself and then searching for that certain     post in the function **GetPostsEndPoint** , the command and the output looks like this:-<br><br>
  ![Capture9](https://user-images.githubusercontent.com/66525380/136668227-8261f999-3458-4c0e-aabd-1f58b82958a4.PNG)<br>
  ![Capture10](https://user-images.githubusercontent.com/66525380/136668235-3cefbce2-e52d-4474-a98c-5ee16b5f6298.PNG)
* List all the posts using the user_id<br>
  Similar to getting a post by its post_id , using the user_id field in the posts collection we can get all the posts made by a     single user by equalizing the user_id that we get from parsing the url in the GET command in the **GetUserPostsEndPoint**, the     command and the output looks like this:-<br>
  ![Capture11](https://user-images.githubusercontent.com/66525380/136668453-7805d7fc-88d9-4bad-8c57-5f5600ef2023.PNG)<br>
  ![Capture12](https://user-images.githubusercontent.com/66525380/136668515-daef1b93-496d-416f-be86-a7419fc4e7c7.PNG)<br>
* Dependencies with the project<br>
  Following the guidelines in the project document I have only used the standard library and the mongodb library for my task, so     to run the project only one dependency must be installed and that could be done by:-<br>
  **go get go.mongodb.org/mongo-driver**
  <br>After the installation the project should be ready to go!!

  

  




  

  

