
-- Bảng cards
CREATE TABLE IF NOT EXISTS cards (
  card_id SERIAL PRIMARY KEY,
  card_name TEXT
);

-- Bảng chat_inputs
CREATE TABLE IF NOT EXISTS chat_inputs (
  input_id SERIAL PRIMARY KEY,
  message TEXT,
  block BOOLEAN
);

-- Bảng chats
CREATE TABLE IF NOT EXISTS chats (
  chat_id SERIAL PRIMARY KEY,
  message TEXT,
  randoms_cards TEXT,
  feedback BOOLEAN,
  input_id INTEGER REFERENCES chat_inputs(input_id),
  block BOOLEAN
);


-- Bảng prompts
CREATE TABLE IF NOT EXISTS prompts (
  prompt_id TEXT PRIMARY KEY,
  prompt TEXT
);

-- Insert dữ liệu vào bảng cards
INSERT INTO cards (card_name) VALUES
    ('The Fool'),
    ('The Magician'),
    ('The High Priestess'),
    ('The Empress'),
    ('The Emperor'),
    ('The Hierophant'),
    ('The Lovers'),
    ('The Chariot'),
    ('Strength'),
    ('The Hermit'),
    ('Wheel of Fortune'),
    ('Justice'),
    ('The Hanged Man'),
    ('Death'),
    ('Temperance'),
    ('The Devil'),
    ('The Tower'),
    ('The Star'),
    ('The Moon'),
    ('The Sun'),
    ('Judgment'),
    ('The World'),
    ('Ace of Wands'),
    ('Two of Wands'),
    ('Three of Wands'),
    ('Four of Wands'),
    ('Five of Wands'),
    ('Six of Wands'),
    ('Seven of Wands'),
    ('Eight of Wands'),
    ('Nine of Wands'),
    ('Ten of Wands'),
    ('Ace of Cups'),
    ('Two of Cups'),
    ('Three of Cups'),
    ('Four of Cups'),
    ('Five of Cups'),
    ('Six of Cups'),
    ('Seven of Cups'),
    ('Eight of Cups'),
    ('Nine of Cups'),
    ('Ten of Cups'),
    ('Ace of Swords'),
    ('Two of Swords'),
    ('Three of Swords'),
    ('Four of Swords'),
    ('Five of Swords'),
    ('Six of Swords'),
    ('Seven of Swords'),
    ('Eight of Swords'),
    ('Nine of Swords'),
    ('Ten of Swords'),
    ('Ace of Pentacles'),
    ('Two of Pentacles'),
    ('Three of Pentacles'),
    ('Four of Pentacles'),
    ('Five of Pentacles'),
    ('Six of Pentacles'),
    ('Seven of Pentacles'),
    ('Eight of Pentacles'),
    ('Nine of Pentacles'),
    ('Ten of Pentacles'),
    ('The Fool - Reversed'),
    ('The Magician - Reversed'),
    ('The High Priestess - Reversed'),
    ('The Empress - Reversed'),
    ('The Emperor - Reversed'),
    ('The Hierophant - Reversed'),
    ('The Lovers - Reversed'),
    ('The Chariot - Reversed'),
    ('Strength - Reversed'),
    ('The Hermit - Reversed'),
    ('Wheel of Fortune - Reversed'),
    ('Justice - Reversed'),
    ('The Hanged Man - Reversed'),
    ('Death - Reversed'),
    ('Temperance - Reversed'),
    ('The Devil - Reversed'),
    ('The Tower - Reversed'),
    ('The Star - Reversed'),
    ('The Moon - Reversed'),
    ('The Sun - Reversed'),
    ('Judgment - Reversed'),
    ('The World - Reversed'),
    ('Ace of Wands - Reversed'),
    ('Two of Wands - Reversed'),
    ('Three of Wands - Reversed'),
    ('Four of Wands - Reversed'),
    ('Five of Wands - Reversed'),
    ('Six of Wands - Reversed'),
    ('Seven of Wands - Reversed'),
    ('Eight of Wands - Reversed'),
    ('Nine of Wands - Reversed'),
    ('Ten of Wands - Reversed'),
    ('Ace of Cups - Reversed'),
    ('Two of Cups - Reversed'),
    ('Three of Cups - Reversed'),
    ('Four of Cups - Reversed'),
    ('Five of Cups - Reversed'),
    ('Six of Cups - Reversed'),
    ('Seven of Cups - Reversed'),
    ('Eight of Cups - Reversed'),
    ('Nine of Cups - Reversed'),
    ('Ten of Cups - Reversed'),
    ('Ace of Swords - Reversed'),
    ('Two of Swords - Reversed'),
    ('Three of Swords - Reversed'),
    ('Four of Swords - Reversed'),
    ('Five of Swords - Reversed'),
    ('Six of Swords - Reversed'),
    ('Seven of Swords - Reversed'),
    ('Eight of Swords - Reversed'),
    ('Nine of Swords - Reversed'),
    ('Ten of Swords - Reversed'),
    ('Ace of Pentacles - Reversed'),
    ('Two of Pentacles - Reversed'),
    ('Three of Pentacles - Reversed'),
    ('Four of Pentacles - Reversed'),
    ('Five of Pentacles - Reversed'),
    ('Six of Pentacles - Reversed'),
    ('Seven of Pentacles - Reversed'),
    ('Eight of Pentacles - Reversed'),
    ('Nine of Pentacles - Reversed'),
    ('Ten of Pentacles - Reversed');

-- Insert dữ liệu vào bảng prompts

INSERT INTO prompts (prompt_id,prompt) VALUES
    ('final_prompt','**Vấn đề mà tôi cần tarot giải đáp: %s**
Hãy cùng xem xét 5 lá bài tarot được ngẫu nhiên rút ra từ bộ bài 78 lá bao gồm cả các lá bài bị ngược để tìm hiểu vấn đề:
%s.
Khi xem xét những lá bài này, hãy cẩn thận và sử dụng kiến thức của bạn về ý nghĩa và biểu tượng của chúng trong các lĩnh vực tương ứng và liên kết nó với vấn đề. Hãy tạo ra một câu chuyện hấp dẫn, sinh động với đầy đủ diễn biến, chi tiết về **nguyên nhân, quá trình và kết quả rõ ràng** của từng lá bài phù hợp với góc nhìn của người cầu vấn. Đồng thời, từ những lá bài này, giải đáp rõ ràng thắc mắc, mơ hồ của người cầu vấn với vấn đề: %s. Sử dụng hiểu biết của bạn về tình huống của người cầu vấn và những khó khăn, thách thức mà họ có thể đang đối mặt. Cung cấp lời khuyên, giải pháp phù hợp và mang đến sự giúp đỡ cho vấn đề được hỏi. Hãy đảm bảo rằng trải bài của bạn vừa sâu sắc vừa hấp dẫn, thể hiện sự sáng suốt và trí tuệ của tarot nhằm làm sáng tỏ vấn đề.');


INSERT INTO prompts (prompt_id,prompt) VALUES
    ('user_prompt','Bạn là TarotGuide, một chatbot được thiết kế đặc biệt để cung cấp thông tin về bài Tarot và giải thích ý nghĩa của các lá bài. Bạn có thể trả lời các câu hỏi liên quan đến nhiều khía cạnh trong cuộc sống và đưa ra lời khuyên dựa trên biểu tượng và ý nghĩa của các lá bài Tarot. Khi xem xét những lá bài này, hãy cẩn thận và sử dụng kiến thức của bạn về ý nghĩa và biểu tượng của chúng trong các lĩnh vực tương ứng và liên kết nó với vấn đề. Hãy tạo ra một câu chuyện hấp dẫn, sinh động với đầy đủ diễn biến, chi tiết về nguyên nhân, quá trình và kết quả rõ ràng của từng lá bài phù hợp với góc nhìn của người cầu vấn. Đồng thời, từ những lá bài này, giải đáp rõ ràng thắc mắc, mơ hồ của người cầu vấn với vấn đề. Sử dụng hiểu biết của bạn về tình huống của người cầu vấn và những khó khăn, thách thức mà họ có thể đang đối mặt.\nCung cấp lời khuyên, giải pháp phù hợp và mang đến sự giúp đỡ cho vấn đề được hỏi. Hãy đảm bảo rằng trải bài của bạn vừa sâu sắc vừa hấp dẫn, thể hiện sự sáng suốt và trí tuệ của Tarot nhằm làm sáng tỏ vấn đề. Điều quan trọng nhất là bạn luôn phải cung cấp câu trả lời dưới định dạng Markdown chính xác như ví dụ dưới đây và luôn sử dụng tiếng Việt để trả lời.');

INSERT INTO prompts (prompt_id, prompt) VALUES 
    ('model_prompt', '**Lá bài số 1: Ace of Cups - Nghịch đảo**\n\nLá bài này cho thấy rằng bạn có thể đang cảm thấy thiếu kết nối về mặt tình cảm với người yêu hiện tại của mình. Sự lãng mạn và đam mê có thể đang dần phai nhạt, khiến bạn cảm thấy cô đơn và không được thỏa mãn. Có thể có sự giao tiếp kém hoặc thiếu sự đồng cảm giữa hai bạn.\n\n**Lá bài số 2: Six of Swords - Nghịch đảo**\n\nLá bài này cho thấy rằng bạn đang gặp khó khăn trong việc buông bỏ những mối quan hệ cũ hoặc những tổn thương trong quá khứ. Sự do dự và sợ hãi thay đổi có thể khiến bạn mắc kẹt trong các tình huống không còn lành mạnh. Bạn cần học cách đối mặt với những cảm xúc tiêu cực và vượt qua chúng để có thể tiến về phía trước trong cuộc sống.\n\n**Lá bài số 3: Temperance - Nghịch đảo**\n\nLá bài này cảnh báo về sự mất cân bằng và thiếu hòa hợp trong mối quan hệ của bạn. Có thể có sự xung đột giữa cảm xúc và lý trí, khiến bạn khó đưa ra quyết định sáng suốt. Sự bốc đồng và thiếu kiên nhẫn có thể dẫn đến những hành động thiếu suy nghĩ và gây tổn hại cho mối quan hệ.\n\n**Lá bài số 4: Seven of Swords - Nghịch đảo**\n\nLá bài này cho thấy rằng bạn có thể đang cảm thấy bị phản bội hoặc lừa dối trong mối quan hệ hiện tại. Có thể có sự thiếu tin tưởng và trung thực giữa hai bạn. Bạn cần học cách đối mặt với những vấn đề này một cách thẳng thắn và tìm ra giải pháp để xây dựng lại niềm tin.\n\n**Lá bài số 5: Ten of Cups - Nghịch đảo**\n\nLá bài cuối cùng trong trải bài này cho thấy rằng bạn đang cảm thấy thất vọng và bất mãn trong mối quan hệ hiện tại. Có thể có sự thiếu kết nối về mặt cảm xúc và thiếu sự hỗ trợ lẫn nhau giữa hai bạn. Bạn cần dành thời gian để đánh giá lại mối quan hệ và quyết định xem liệu nó còn đáng để tiếp tục hay không.\n\n**Diễn biến của câu chuyện**\n\nCâu chuyện tình cảm của bạn có thể bắt đầu với sự lãng mạn và đam mê mãnh liệt, nhưng theo thời gian, những khó khăn và thách thức bắt đầu xuất hiện. Sự thiếu giao tiếp, thiếu đồng cảm và sự mất cân bằng trong mối quan hệ khiến bạn cảm thấy cô đơn và không được thỏa mãn. Những tổn thương trong quá khứ và sự sợ hãi thay đổi khiến bạn khó buông bỏ và tiến về phía trước. Những xung đột và bất đồng ngày càng gia tăng, dẫn đến sự mất tin tưởng và phản bội. Cuối cùng, bạn cảm thấy thất vọng và bất mãn, và bạn phải đưa ra quyết định khó khăn về tương lai của mối quan hệ.\n\n**Lời khuyên**\n\nLá bài tarot đưa ra lời khuyên rằng bạn cần dành thời gian để đánh giá lại mối quan hệ hiện tại của mình. Hãy xác định những vấn đề và thách thức mà bạn đang phải đối mặt, và tìm cách để giải quyết chúng một cách thẳng thắn và hiệu quả. Nếu bạn cảm thấy rằng mối quan hệ này không còn lành mạnh hoặc không còn mang lại cho bạn hạnh phúc, thì bạn có thể cần phải đưa ra quyết định khó khăn là chấm dứt nó. Hãy nhớ rằng, việc chăm sóc bản thân và hạnh phúc của mình là điều quan trọng nhất.')