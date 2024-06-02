package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

	posts := []string{
		`From TCS → SDE-1 Flipkart
.
.
Recently, A dear friend of mine secured an SDE-1 position at Flipkart, and here’s her experience → 

This started with the Shortlisting round after clearing it there was the 

📍1st Round - Machine Coding, Duration: 2 Hours.
The machine coding round was on a Google Meet, 30-35 people were there. We all got the same question, 

An SDE-3 from Flipkart explained the question in the first 30 mins. 

In the remaining 1:30 Hour we had to write the code. I was able to complete the mandatory SDE 1 feature with a bonus.

"After a few days, I got a call for the PSDS round."

📍2nd Round- Problem Solving & DSA (1 Hour)

I was given two Leetcode (Hard-level) questions. I was able to solve it in 30 minutes, it was based on Array, Monotonic Stack. 

The other question was a bit easier but I struggled a little on this one, I explained the approach and wrote a code. It was not exactly correct. 

On the same evening, I got a call for the Final round. It was with the Hiring Manager

📍3rd Round- Hiring Manager 
We had a discussion about my past work, a few behavioural questions & one coding problem was given to understand my approach only.

After 3 days I got the good news that I was selected for the SDE-1 Role at Flipkart 

This was all possible because of her consistency and dedication. 

If you’re someone who is aiming for product-based companies and requires some help in the preparation I would suggest you check Bosscoder Academy 

Check here: https://bit.ly/3wXu3BS 

They can help you with
✅ Structured curriculum to learn problem-solving in DSA, System Design, and Full stack development.
✅ Personal Mentorship from industry experts.
✅ Live Classes & 24/7 Doubt Support.
✅ Real-life, relevant projects with Placement Assistance

#tcs #flipkart #consistency #softwareengineer`,
		`While making a switch from Service ➡️ Product company, most Software Engineers get overwhelmed.
.
.
And feeling demotivated is normal. You should embrace the process of preparing. 

In product-based companies, where skills like DSA, System Design, etc. often lead us to imitate others. But in this world, each person is unique. 👍

Everyone has their own pace of learning So, instead of trying to copy someone else's path, focus on understanding your strengths and weaknesses. 


Also, don't underestimate the importance of soft skills. 

Communication, problem-solving, and teamwork are equally important to crack these interviews at product-based companies. 

So, make sure to work on these skills alongside technical expertise. 🙌

Lastly, maintain a Positive Mindset. Job hunting can be challenging, but every setback is an opportunity to learn and grow. Your patience will pay off. 

After all this, If you need any assistance in your preparation journey I would suggest you check out Bosscoder Academy

Check them here: https://bit.ly/4dXYOYn 

With them, you get:
✅ Structured programs for learning DSA, System Design, and Full Stack Development.
✅ Personal mentorship from industry veterans.
✅ Live classes and 24/7 doubt support.
`,
		`Before cracking Flipkart, I got rejected by a few companies like Oracle, etc.! 😓

After 2 consecutive rejections,, demotivated me greatly. 

But instead of sulking, I used them as a motivation, and this became a turning point. 

Here are 2 major mistakes that I was making: 

📍Before the interview, my basics were not clear and I was not prepared well.

📍While doing DSA, I just focused on solving 100s of problems, instead of focusing on the variety of it. 

Because of this, my solutions were not optimized enough and I was rejected. 

However, I was able to turn everything around, with my Consistency & hard work. 🙌

No, matter what, I kept on improving myself. Even though I had limited time, I made sure to utilize it well. 

→ I prepared a daily schedule every night & tried to stick with it. 
→ Setting my priorities helped me be more accountable for my learning. 💪

These things helped me finally crack Flipkart! ❤️

However, I understand, that not every working professional has that much time to prepare properly and be consistent.

But if you aspire to prepare strategically, and fully for your next technical interview, I’d suggest taking help from industry veterans, like Bosscoder Academy. 

👉Checkout here- https://bit.ly/3UToduD 

They can help you by providing →

✅ Structured curriculum to learn problem-solving in DSA, System Design, and Full stack development.
✅ Personal Mentorship from industry experts.
✅ Live Classes & 24/7 Doubt Support.
✅ Real-life, relevant projects with Placement Assistance
Work hard for you dream. Hustle a little! 🔥




hashtag#rejections hashtag#sponsoredpost hashtag#flipkart`,
		`Before joining Flipkart, I too had to go through a phase where I was rejected due to DSA. 😪

My poor problem-solving skills and DSA became a major reason for the setback.

However, to improve myself, I pushed myself hard, I still remember the nights spend revising Trees, DP, etc. 😓
 
Honestly, cracking interviews at any product-based company is not just by luck. 

So, here’s something that I learned in my whole journey to Flipkart→ 

📌 Motivation is good, but make sure it remains constant throughout your preparation. 📝

Before you plan on mastering a thing completely- rate yourself topic by topic. This will help you analyze your weak and strong areas.

So, start with a well-structured plan, follow it through, and be consistent.

📌 Some topics like Graphs, DP, and Trees may appear difficult, but they should not put a break on your progress. 💪🏼

Overcoming these challenging topics is important for your success. Always try to take the help of mentors or peers if you feel stuck at any point.

📌 Make sure to practice daily on Leetcode to build your DSA & problem-solving skills. 

Following these things CONSISTENTLY, will help you grow your skills. 🙌🏽

However, I understand that working professionals find it difficult to do all these things independently.

So, if you’re looking for an end-to-end solution and dream of landing at a top product company, make sure to check Bosscoder Academy.

📍𝗖𝗵𝗲𝗰𝗸 𝘁𝗵𝗲𝗺 𝗼𝘂𝘁 𝗵𝗲𝗿𝗲 → https://bit.ly/4dkoC0n 

Join the Success Story of 1000+ Software Engineers with the help of:
✅ Structured curriculum to learn problem-solving in DSA, System Design and Full stack development.
✅ Personal Mentorship from industry experts.
✅ Live Classes & 24/7 Doubt Support.
✅ Real-life, relevant projects with Placement Assistance





hashtag#motivation hashtag#collab hashtag#flipkart
`,
		`Got rejected from my Dream MAANG Company! 💔 
.
In 2021, I faced multiple rejections from top product-based companies. 

Each "NO" reminded me of the gaps I needed to fill, and instead of giving up, I focused on building my skills.

Starting with DSA → I built a strong hold on the basic fundamentals and soon I moved to more complex topics. 

After my first rejection, I started to take System Design more seriously and dedicated more time to it. Focused on both LLD and HLD. ✔️

Did a lot of problem-solving – at the start it was a little difficult and I was not fast, but as I started solving more problems day by day I started to gain more confidence and became efficient at it. 💪

Not only this, I gave a lot of mock interviews, which helped me assess my mistakes and work on them.

Slowly, I became more confident. Fast forward to 2022, I received an offer from a MAANG Company 🎉

The journey was tough, but it taught me the importance of continuous learning. 

And if you’re someone who is going through a similar phase and needs a little help, I would suggest you check Bosscoder Academy

Check here: https://bit.ly/4bzPBDL 
 
They can help you with
✅ Structured curriculum to learn problem-solving in DSA, System Design, and Full stack development.
✅ Personal Mentorship from industry experts.
✅ Live Classes & 24/7 Doubt Support.
✅ Real-life, relevant projects with Placement Assistance`,
		// Include other posts here
	}

	prompt := "Generate a post inspired by the following posts:\n\n"
	for _, post := range posts {
		prompt += post + "\n\n"
	}
	prompt += "Stick to the format of the posts like transition from X service based company to product based and stuff like that. Use similar emojis as the main post and pointers. Mention stuff like 7LPA - 42LPA and 8LPA- 80LPA and also different service based company names like Wipro TCS and product based like AMazon MSFT and others"

	requestBody, err := json.Marshal(OpenAIRequest{
		Model:       "gpt-3.5-turbo",
		Messages:    []Message{{Role: "user", Content: prompt}},
		Temperature: 0.7,
	})
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed with status code: %d", resp.StatusCode)
	}

	var openAIResponse OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResponse); err != nil {
		log.Fatalf("Failed to decode response body: %v", err)
	}

	if len(openAIResponse.Choices) > 0 {
		responseText := openAIResponse.Choices[0].Message.Content
		fmt.Println("Response Content:", responseText)

		// Add emojis at the end of each keyword in the response
		responseTextWithEmojis := addEmojis(responseText)

		// Save response to a file in text format
		if err := saveToFile("response.txt", responseTextWithEmojis); err != nil {
			log.Fatalf("Failed to save response to file: %v", err)
		}
	} else {
		log.Println("No choices returned from OpenAI")
	}
}

func addEmojis(text string) string {
	// Add emojis at the end of each keyword in the text
	emojis := map[string]string{
		"product-based company": " 💼",
		"consistency":           " 📈",
		"dedication":            " 💪",
		"mentorship programs":   " 👨‍🏫",
	}

	// Replace specific keywords with emojis
	for keyword, emoji := range emojis {
		text = strings.ReplaceAll(text, keyword, keyword+emoji)
	}

	return text
}

func saveToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Println("Response saved to", filename)
	return nil
}
