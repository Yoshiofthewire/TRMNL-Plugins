# How to Set Up the Daily Comic from The Far Side

<kbd>![daily-comic-rankings-plugin](https://github.com/Yoshiofthewire/TRMNL-Plugins/blob/main/FarSide.png)</kbd>

## Step 0: [Optional] Create a private scraping server
As an Optional step, create your own private scraping server.

---
1. Install Go: Make sure you have Go installed (https://golang.org/doc/install).
2. Download main.go and place in a perminant location
3. Initalize the project: go mod init farside_go
4. Install goquery: go get github.com/PuerkitoBio/goquery; go mod tidy
5. Run the Scrapper: go run main.go
6. Make this server publicly avablible on port 8123
---

## Step 1: Create a New Private Plugin
Log in to your TRMNL dashboard.
On the left-hand menu, click on the 'Go to Plugins' button.
Find the 'Private Plugin' Plugin to create a Private Plugin.
Click 'Add new' to create a new Private Plugin.

## Step 2: Set up the Polling Strategy
Name your plugin (e.g., "Daily Comic") then scroll down to the Strategy section.
Choose the Polling strategy from the Strategy dropdown menu.
In the Polling URL field, enter this URL (Or the Server Created in Stage 0):

```
https://farapi.urlxl.com/
```

Click Save. Once it is saved, the 'Edit Markup' button is now available.

## Step 3: Add the HTML Markup
Click the 'Edit Markup' button.

Copy and paste the following code into the Markup box. This code will display a daily comic supplied from xkcd.
```
<div class="view comic-view bg-white">
  <div class="layout layout--col gap--space-between">
    <div class="columns">
      <div class="column">
        <div class="content content" style="max-height: 95hv">
          <!-- Centered Comic Image -->
          <img src="{{ img }}" alt="{{ title }}" style="display: block; margin: 0 auto; max-width: 80%; height: 350px;  border-radius: 5px;" />
          <p style="text-align:center">{{ title }}"</p>
        </div>
      </div>
    </div>
  </div>
  <div class="title_bar">
    <!-- Left-Aligned Title Bar Content -->
    <img class="image" src="https://siteassets.thefarside.com/packs/media/images/brand/logo/tfs_logo-cb7a52999a9300e203e36865dc428f71.svg" alt="The Far Side by Gary Larson" />

  </div>
</div>


```

## Step 4: Save and Activate the Plugin
Once you have entered the markup, click Save to store the plugin.
Navigate to the Playlists tab in your TRMNL dashboard.
Drag and drop your new Daily Comic plugin to the top of your playlist if not automatically added.

## Step 5: View the Daily Comic on Your Device
Once refreshed, your TRMNL device will display the Daily Comic.
