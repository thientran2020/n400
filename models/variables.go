package models

const (
	LINK            = "uscis.gov/citizenship/testupdates"
	PROMT           = "Enter to continue. [Q] to quit: "
	WELCOME_MESSAGE = "\n_________WELCOME TO PREPATION N-400 TEST_________\n\n"
	EXIT_MESSAGE    = "\nGOOD LUCK WITH YOUR EXAM ^.^\n"
	RED             = "\033[1;31m"
	BLUE            = "\033[1;34m"
	GREEN           = "\033[1;32m"
	YELLOW          = "\033[1;33m"
	UWHITE          = "\033[4;37m"
	COLOROFF        = "\033[0m"
	MARK            = "\u2713"
)

var (
	CATEGORY_KEYS    = []int{1, 58, 88}
	SUBCATEGORY_KEYS = []int{1, 13, 48, 58, 71, 78, 88, 96, 99}
)

var CATEGORIES = map[int]string{
	1:  "American Government",
	58: "American History",
	88: "Integrated Civics",
}

var SUBCATEGORIES = map[int]string{
	1:  "Principles of American Democracy",
	13: "System of Government",
	48: "Right and Responsibilities",
	58: "Colonial Period and Independence",
	71: "1800's",
	78: "Recent American History and Other Important Historical Information",
	88: "Geography",
	96: "Symbols",
	99: "Holidays",
}

var QUESTIONS = []string{
	"What is the supreme law of the land?",
	"What does the Constitution do?",
	"The idea of self-government is in the first THREE words of the Constitution. What are these words?",
	"What is an amendment?",
	"What do we call the first 10 amendments to the Constitution?",
	"What is ONE right or freedom from the First Amendment?",
	"How many amendments does the Constitution have?",
	"What did the Declaration of Independence do?",
	"What are TWO rights in the Declaration of Independence?",
	"What os freedom of religion?",
	"What is the economic system in the U.S.?",
	"What is the 'rule of law'?",
	// -------------------------- ///
	"Name ONE branch or part of the government?",
	"What stops ONE branch of government from becoming too powerful?",
	"Who is in charge of the executive branch?",
	"Who makes federal laws?",
	"What are the TWO parts of the U.S. Congress?",
	"How many U.S. Senators are there?",
	"We elect a U.S. Senator for how many years?",
	"Who is ONE of your state's U.S. Senators now? Dianne Feinstein || Alex Padilla", // ***
	"The House of Representatives has how many voting members?",
	"We elect a U.S. Representative for how many years?",
	"Name your U.S. Representative. Ro Khanna", // ***
	"Who does a U.S. Senator present?",
	"Why do some states have more Representatives than other states?",
	"We elect a President for how many years?",
	"In what month do we vote for President?",
	"What is the name of the President of the U.S. now? Joe Biden",          // ***
	"What is the name of the Vice President of the U.S. now? Kamala Harris", // ***
	"If the President can no longer serve, who becomes President?",
	"If both the President and the Vice President can no longer serve, who becomes President?",
	"Who is the Commander in Chief of the military?",
	"Who signs bills to become laws?",
	"Who vetoes bills?",
	"What does the President's Cabinet do?",
	"What are TWO Cabinet-level positions?",
	"What does the judicial branch do?",
	"What is the highest court in the U.S.?",
	"How many justices are on the Supreme Court?",
	"Who is the Chief Justice of the U.S. now? John Roberts", // ***
	"Under our Constitution, some powers belong to the federal government. What is ONE power of the federal government?",
	"Under our Constitution, some powers belong to the states. What is ONE power of the states?",
	"Who is the Governor of your state now? Gavin Newsom", // ***
	"What is the capital of your state? Sacramento",       // ***
	"What are the TWO major political parties in the U.S.?",
	"What is the political party of the President now? Democratic Party",                // ***
	"What is the name of the Speaker of the House of Representatives now? Nancy Pelosi", // ***
	// -------------------------- ///
	"There are 4 amendments to the Constitution about who can vote. Describe ONE of them.",
	"What is ONE responsibility that is only for U.S. citizens?",
	"Name ONE right only for U.S. citizens.",
	"What are TWO rights of everyone living in the U.S.?",
	"What do we show loyalty to when we say the Pledge of Allegiance?",
	"What is one promise you make when you become a U.S. citizen?",
	"How old do citizens have to be to vote for President?",
	"What are TWO ways that Americans can participate in their democracy?",
	"When is the last day you can send in federal income tax forms?",
	"When must all men register for the Selective Service?",
	// -------------------------- ///
	"What is ONE reason colonists came to America?",
	"Who lived in America before the Europeans arrived?",
	"What group of people was taken to America and sold as slaves?",
	"Why did the colonists fight the British?",
	"Who wrote the Declaration of Independence?",
	"When was the Declaration of Independence adopted?",
	"There were 13 original states. Name THREE?",
	"What happened at the Consitutional Convention?",
	"When was the Constitution written?",
	"The Federalist Papers supported the passage of the U.S. Constitution. Name ONE of the writers.",
	"What is ONE thing Benjamin Franklin is famous for?",
	"Who is the 'Father of Our Country'?",
	"Who was the first President?",
	// -------------------------- ///
	"What territory did the U.S. buy from France in 1803?",
	"Name ONE war fought by the U.S. in the 1800s.",
	"Name the U.S. war between the North and the South.",
	"Name ONE problem that led to the Civil War.",
	"What was ONE important thing that Abraham Lincoln did?",
	"What did the Emancipation Proclamation do?",
	"What did Susan B. Anthony do?",
	// -------------------------- ///
	"Name ONE war fought by the U.S. in the 1900s.",
	"Who was President during World War I?",
	"Who was President during the Great Depression and World War II?",
	"Who did the U.S. fight in World War II?",
	"Before he was President, Eisenhower was a general. What war he was in?",
	"During the Cold War, what was the main concern of the U.S.?",
	"What movement tried to end racial discrimination?",
	"What did Martin Luther King, Jr. do?",
	"What major event happend on September 11, 2001, in the U.S.?",
	"Name ONE American Indian tribe in the U.S.",
	// -------------------------- ///
	"Name ONE of the TWO longest rivers in the U.S.",
	"What ocean is on the West Coast of the U.S.?",
	"What ocean is on the East Coast of the U.S.?",
	"Name ONE U.S. territory",
	"Name ONE state that borders Canada.",
	"Name ONE state that borders Mexico.",
	"What is the capital of the U.S.?",
	"Where is the Statue of Liberty?",
	// -------------------------- ///
	"Why does the flag have 13 stripes?",
	"Why does the flag have 50 stars?",
	"What is the name of the national anthem?",
	// -------------------------- ///
	"When do we celebrate Independence Day?",
	"Name TWO national U.S. holidays?",
}
