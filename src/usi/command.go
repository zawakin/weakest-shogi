package usi

// Command is a type of a usi command.
type Command string

// コマンドの説明は、http://shogidokoro.starfree.jp/usi.html より引用。

const (
	// CommandFirstUSI は、GUIがエンジンを起動した時に最初に送るコマンドです。
	// エンジンは、このコマンドを受信したらすぐにidコマンドを返し、必要に応じてoptionコマンドを返します。そして最後にusiokコマンドを返す必要があります。
	CommandFirstUSI Command = "usi"

	// CommandIsReady は、対局開始前に送るコマンドです。
	// エンジンは対局準備ができたらreadyokを返すことになります。
	CommandIsReady Command = "isready"

	// CommandSetOption は、エンジンに対して値を設定する時に送ります。
	CommandSetOption Command = "setoption"

	// CommandNewGame は、対局開始時に送るコマンドです。
	// これで対局開始になります。
	CommandNewGame Command = "usinewgame"

	// CommandSetPosition は、思考開始局面をエンジンに知らせるためのコマンドです。
	// エンジンに思考を開始させる場合、positionコマンドで思考開始局面を送り、それに続けてgoコマンドを送って思考開始を命令することになります。
	CommandSetPosition Command = "position"

	// CommandGo は、思考開始の合図のためのコマンドです。エンジンはこれを受信すると思考を開始します。
	CommandGo Command = "go"

	// CommandStop は、エンジンに対し思考停止を命令するコマンドです。
	// エンジンはこれを受信したら、できるだけすぐ思考を中断し、bestmoveで指し手を返す必要があります。（現時点で最善と考えている手を返すようにして下さい。）
	CommandStop Command = "stop"

	// CommandPonderhit は、エンジンが先読み中、前回のbestmoveコマンドでエンジンが予想した通りの手を相手が指した時に送ります。
	// エンジンはこれを受信すると、先読み思考から通常の思考に切り替わることになり、任意の時点でbestmoveで指し手を返すことができます。
	CommandPonderhit Command = "ponderhit"

	// CommandQuit は、アプリケーション終了を命令するコマンドです。エンジンはこれを受信したらすぐに終了する必要があります。
	CommandQuit Command = "quit"

	// CommandGameOver は、対局終了をエンジンに知らせるコマンドです。
	// gameoverのあと、エンジンの結果に応じてwin, lose, drawのいずれかのパラメータも一緒に送られます。
	// エンジンはgameoverを受信したら対局状態を終了して、対局待ち状態になります。その後、isready及びusinewgameを受信すると次の対局開始ということになります。
	CommandGameOver Command = "gameover"
)
