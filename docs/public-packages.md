<!-- insert:PUBLIC_PACKAGES -->
# Public Packages

- [**`github.com/zekroTJA/shinpuru/pkg/validators`**](/pkg/validators)  
  *Package validators provides some (more or less) general purpose validator functions for user inputs.*

- [**`github.com/zekroTJA/shinpuru/pkg/checksum`**](/pkg/checksum)  
  *Package checksum provides functions to generate a hash sum from any given object.*

- [**`github.com/zekroTJA/shinpuru/pkg/stringutil`**](/pkg/stringutil)  
  *Package stringutil provides generl string utility functions.*

- [**`github.com/zekroTJA/shinpuru/pkg/thumbnail`**](/pkg/thumbnail)  
  *Package thumbnail provides simple functionalities to generate thumbnails from images with a max witdh or height.*

- [**`github.com/zekroTJA/shinpuru/pkg/multierror`**](/pkg/multierror)  
  *Package multierror impements handling multiple errors as one error object.*

- [**`github.com/zekroTJA/shinpuru/pkg/jdoodle`**](/pkg/jdoodle)  
  *Package jdoodle provides an API wrapper for the jdoodle execute and credit-spent REST API.*

- [**`github.com/zekroTJA/shinpuru/pkg/lctimer`**](/pkg/lctimer)  
  *Package lctimer provides a life cycle timer which calls registered callback handlers on timer elapse. This package is a huge buggy piece of crap, please don't use it. :)*

- [**`github.com/zekroTJA/shinpuru/pkg/rediscmdstore`**](/pkg/rediscmdstore)  
  *Package rediscmdstore provides an implementation of github.com/zekrotja/ken/store.CommandStore using a redis client to store the command cache.*

- [**`github.com/zekroTJA/shinpuru/pkg/etag`**](/pkg/etag)  
  *Package etag implements generation functionalities for the ETag specification of RFC7273 2.3. https://tools.ietf.org/html/rfc7232#section-2.3.1*

- [**`github.com/zekroTJA/shinpuru/pkg/fetch`**](/pkg/fetch)  
  *Package fetch provides functionalities to fetch roles, channels, members and users by so called resolavbles. That means, these functions try to match a member, role or channel by their names, displaynames, IDs or mentions as greedy as prossible.*

- [**`github.com/zekroTJA/shinpuru/pkg/argp`**](/pkg/argp)  
  *Package argp is a stupid simple flag (argument) parser which allows to parse flags without panicing when non-registered flags are passed.*

- [**`github.com/zekroTJA/shinpuru/pkg/inline`**](/pkg/inline)  
  *Package inline provides general inline operation functions like inline if or null coalescence.*

- [**`github.com/zekroTJA/shinpuru/pkg/timerstack`**](/pkg/timerstack)  
  *Package timerstack provides a timer which can execute multiple delayed functions one after one.*

- [**`github.com/zekroTJA/shinpuru/pkg/twitchnotify`**](/pkg/twitchnotify)  
  *Package twitchnotify provides functionalities to watch the state of twitch streams and notifying changes by polling the twitch REST API.*

- [**`github.com/zekroTJA/shinpuru/pkg/boolutil`**](/pkg/boolutil)  
  *Package boolutil provides simple utility functions around booleans.*

- [**`github.com/zekroTJA/shinpuru/pkg/bytecount`**](/pkg/bytecount)  
  *Package bytecount provides functionalities to format byte counts.*

- [**`github.com/zekroTJA/shinpuru/pkg/timeutil`**](/pkg/timeutil)  
  *Package timeutil provides some general purpose functionalities around the time package.*

- [**`github.com/zekroTJA/shinpuru/pkg/httpreq`**](/pkg/httpreq)  
  *Package httpreq provides general utilities for around net/http requests for a simpler API and extra utilities for parsing JSON request and response boddies.*

- [**`github.com/zekroTJA/shinpuru/pkg/voidbuffer`**](/pkg/voidbuffer)  
  *Package voidbuffer provides a simple, concurrency proof push buffer with a fixed size which "removes" firstly pushed values when fully filled.*

- [**`github.com/zekroTJA/shinpuru/pkg/lokiwriter`**](/pkg/lokiwriter)  
  *Package lokiwriter implements rogu.Writer to push logs to a Grafana Loki instance.*

- [**`github.com/zekroTJA/shinpuru/pkg/roleutil`**](/pkg/roleutil)  
  *Package roleutil provides general purpose utilities for discordgo.Role objects and arrays.*

- [**`github.com/zekroTJA/shinpuru/pkg/slices`**](/pkg/slices)  
  *Package slices adds generic utility functionalities for slices.*

- [**`github.com/zekroTJA/shinpuru/pkg/logmsg`**](/pkg/logmsg)  
  *No package description.*

- [**`github.com/zekroTJA/shinpuru/pkg/permissions`**](/pkg/permissions)  
  *Package permissions provides functionalities to calculate, update and merge arrays of permission domain rules. Read this to get more information about how permission domains and rules are working: https://github.com/zekroTJA/shinpuru/wiki/Permissions-Guide*

- [**`github.com/zekroTJA/shinpuru/pkg/hammertime`**](/pkg/hammertime)  
  *Package hammertime provides functionailities to format a time.Time into a Discord timestamp mention. The name was used after the very useful web app hammertime.djdavid98.art.*

- [**`github.com/zekroTJA/shinpuru/pkg/discordutil`**](/pkg/discordutil)  
  *Package discordutil provides general purpose extensuion functionalities for discordgo.*

- [**`github.com/zekroTJA/shinpuru/pkg/onetimeauth`**](/pkg/onetimeauth)  
  *Package onetimeout provides short duration valid JWT tokens which are only valid exactly once.*

- [**`github.com/zekroTJA/shinpuru/pkg/limiter`**](/pkg/limiter)  
  *Package limiter provides a fiber middleware for a bucket based request rate limiter.*

- [**`github.com/zekroTJA/shinpuru/pkg/angularservice`**](/pkg/angularservice)  
  *Package angularservice provides bindings to start an Angular development server via the Angular CLI.*

- [**`github.com/zekroTJA/shinpuru/pkg/regexputil`**](/pkg/regexputil)  
  *Package regexutil provides additional utility functions used with regular expressions.*

- [**`github.com/zekroTJA/shinpuru/pkg/colors`**](/pkg/colors)  
  *Package color provides general utilities for image/color objects and color codes.*

- [**`github.com/zekroTJA/shinpuru/pkg/random`**](/pkg/random)  
  *Package random provides some general purpose cryptographically pseudo-random utilities.*

- [**`github.com/zekroTJA/shinpuru/pkg/versioncheck`**](/pkg/versioncheck)  
  *Package versioncheck provides endpoints to retrieve version information via different providers and utilities to compare versions.*

- [**`github.com/zekroTJA/shinpuru/pkg/embedbuilder`**](/pkg/embedbuilder)  
  *Package embedbuilder provides a builder pattern to create discordgo message embeds.*

- [**`github.com/zekroTJA/shinpuru/pkg/hashutil`**](/pkg/hashutil)  
  *Package hashutil provides general utility functionalities to generate simple and fast hashes with salt and pepper.*

- [**`github.com/zekroTJA/shinpuru/pkg/ctypes`**](/pkg/ctypes)  
  *Package ctype provides some custom types with useful function extensions.*

- [**`github.com/zekroTJA/shinpuru/pkg/msgcollector`**](/pkg/msgcollector)  
  *Package msgcollector provides functionalities to collect messages in a channel in conect of a single command request.*

- [**`github.com/zekroTJA/shinpuru/pkg/acceptmsg`**](/pkg/acceptmsg)  
  *Package acceptmsg provides a message model for discordgo which can be accepted or declined via message reactions.*

- [**`github.com/zekroTJA/shinpuru/pkg/startuptime`**](/pkg/startuptime)  
  *Package startuptime provides simple functionalities to measure the startup time of an application.*

- [**`github.com/zekroTJA/shinpuru/pkg/discordoauth`**](/pkg/discordoauth)  
  *package discordoauth provides fasthttp handlers to authenticate with via the Discord OAuth2 endpoint.*

- [**`github.com/zekroTJA/shinpuru/pkg/giphy`**](/pkg/giphy)  
  *Package giphy provides a crappy and inclomplete - but at least bloat free - Giphy API client.*

- [**`github.com/zekroTJA/shinpuru/pkg/mody`**](/pkg/mody)  
  *Package mody allows to modify fields in an object.*

- [**`github.com/zekroTJA/shinpuru/pkg/intutil`**](/pkg/intutil)  
  *Package intutil provides some utility functionalities for integers.*

- [**`github.com/zekroTJA/shinpuru/pkg/mimefix`**](/pkg/mimefix)  
  *+build windows*

