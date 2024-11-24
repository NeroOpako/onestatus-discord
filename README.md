# OneStatus - Discord Integration

OneStatus is a Go application that retrieves your Discord activity and updates your status on OneStatus.

## Usage

Follow these steps:

1. **Create a Discord Bot**:  
   - Visit the [Discord Developer Portal](https://discord.com/developers/applications).  
   - Create a new application and enable the "Presence Intent."  
   - Retrieve the bot's token and application ID.

2. **Add the Bot to a Server**:  
   - Invite your bot to a new private Discord server where it can monitor your activity.   

3. **Create an app password**:  
   - Create an app password in the Bluesky app.  

4. **Configure the Project**:  
   - Create a `secrets.json` file in the project directory and add the following details:  
     ```json
     {
       "discord_app_token": "<Your Discord Token>",
       "discord_app_id": "<Your Discord Bot Application ID>",
       "bsky_user_server": "https://bsky.social/",
       "bsky_user_name": "<Your Bluesky Username>",
       "bsky_user_password": "<Your App Password>"
     }
     ```

5. **Run the Project**:  
   Execute the following command to start the app:  
   ```bash
   go run .
